package models

import (

	"time"
)

type Iplist struct {
	*Model
	Ip          string `json:"ip"`
	Port        string `json:"port"`
	Protocol    string `json:"protocol"`
	Cms         string `json:"cms"`
	Language    string `json:"language"`
	Portnum     int    `json:"portnum"`
	Url         string `json:"url"`
	Loginurl    string `json:"loginurl"`
	Title       string `json:"title"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

func GetIplistBrute(port int, protocol string) (iplist []Iplist) {
	db.Select("ip,port").Where("port = ?", port).Or("protocol = ?", protocol).Find(&iplist)
	return
}

func GetIplist(pageNum int, pageSize int, maps interface{}) (iplist []Iplist) {
	dbTmp := db
	querys := maps.(map[string]interface{})

	if querys["protocol"] != nil {
		dbTmp = dbTmp.Where("protocol LIKE ?", "%"+querys["protocol"].(string)+"%")
	}


	if querys["ip"] != nil {
		dbTmp = dbTmp.Where("ip LIKE ?", "%"+querys["ip"].(string)+"%")
	}

	if querys["port"] != nil {
		dbTmp = dbTmp.Where("port LIKE ?", "%"+querys["port"].(string)+"%")
	}

	if querys["title"] != nil {
		dbTmp = dbTmp.Where("title LIKE ?", "%"+querys["title"].(string)+"%")
	}

	if querys["finger"] != nil {
		dbTmp = dbTmp.Where("cms LIKE ?", "%"+querys["finger"].(string)+"%")
	}

	dbTmp.Offset(pageNum).Limit(pageSize).Order("updated_time  desc").Find(&iplist)
	return
}


func GetIplistTotal(maps interface{}) (count int) {
	dbTmp := db
	querys := maps.(map[string]interface{})
	if querys["protocol"] != nil {
		dbTmp = dbTmp.Where("protocol LIKE ?", "%"+querys["protocol"].(string)+"%")
	}

	if querys["ip"] != nil {
		dbTmp = dbTmp.Where("ip LIKE ?", "%"+querys["ip"].(string)+"%")
	}

	if querys["port"] != nil {
		dbTmp = dbTmp.Where("port LIKE ?", "%"+querys["port"].(string)+"%")
	}

	if querys["title"] != nil {
		dbTmp = dbTmp.Where("title LIKE ?", "%"+querys["title"].(string)+"%")
	}

	if querys["finger"] != nil {
		dbTmp = dbTmp.Where("cms LIKE ?", "%"+querys["finger"].(string)+"%")
	}

	dbTmp.Model(&Iplist{}).Count(&count)
	return
}


func GetIplistTotalDistinct() (count int) {
	db.Model(&Iplist{}).Select("count(distinct(ip))").Count(&count)
	return
}



func ExistIplist(ip, port string) (bool, int) {
	var iplist Iplist
	db.Select("id").Where("ip = ? and port = ? ", ip, port).First(&iplist)
	//如果返回的id>0，也就是数据库里存在过了数据
	if iplist.ID > 0 {
		return true, iplist.ID
	}

	return false, iplist.ID
}

func EditIplist(id int, data interface{}) bool {
	db.Model(&Iplist{}).Where("id = ?", id).Updates(data)
	return true
}

func EditIplistByIp(ip ,port string, data interface{}) bool {
	db.Model(&Iplist{}).Where("ip = ? AND port = ?", ip, port).Updates(data)
	return true
}

func EditIplistByUrl(loginurl string, data interface{}) bool {
	db.Model(&Iplist{}).Where("loginurl = ? ", loginurl).Updates(data)
	return true
}


//创建任务，返回任务id
func AddIplist(data map[string]interface{}) {
	nowTime := time.Now().Format("20060102150405")
	iplist := Iplist{
		Ip:          data["ip"].(string),
		Port:        data["port"].(string),
		Protocol:    data["protocol"].(string),
		CreatedTime: nowTime,
		UpdatedTime: nowTime,
	}
	db.Create(&iplist)
}


func GetIplistTitle() (iplist []Iplist) {
	db.Find(&iplist)
	return
}

// 获取全部http资产
func GetIplistHttp() (iplist []Iplist) {
	db.Where("protocol = ?","http").Find(&iplist)
	return
}


//根据日期删除
//DELETE FROM `iplist`  WHERE updated_time  < '2020-12-27 11:47:56';
func DelIplistUpdate(update string) (iplist []Iplist) {
	db.Where("updated_time  < '"+update+"'").Delete(&iplist)
	return
}


//获取柱状图结果 SELECT port as port, COUNT(*) as portnum FROM `iplist`   GROUP BY port ORDER BY portnum desc;
func GetIplistGroupBy() (iplist []Iplist) {
	var port string
	var portnum int
	rows, _ := db.Model(&Iplist{}).Select("port as port, COUNT(*) as portnum").Group("port").Order("portnum desc").Limit(20).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&port, &portnum)
		iplistTmp := Iplist{
			Port:    port,
			Portnum:    portnum,
		}
		iplist = append(iplist,iplistTmp)
	}

	return
}

//获取柱状图结果 SELECT port as port, COUNT(*) as portnum FROM `iplist`   GROUP BY port ORDER BY portnum desc;
func GetIplistGroupByProtocol() (iplist []Iplist) {
	var protocol string
	var portnum int

	rows, _ := db.Model(&Iplist{}).Select("protocol as protocol, COUNT(*) as portnum").Group("protocol").Order("portnum desc").Limit(20).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&protocol, &portnum)
		iplistTmp := Iplist{
			Protocol:    protocol,
			Portnum:    portnum,
		}
		iplist = append(iplist,iplistTmp)
	}
	return
}


//柱状图 近7日新增资产
func GetIplistGroupByDay(updateTime string) (count int) {
	db.Model(&Iplist{}).Where("updated_time LIKE ?", "%"+updateTime+"%").Count(&count)
	return
}


