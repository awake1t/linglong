package models

import (
	"strconv"
	"time"
)

type Portbruteres struct {
	*Model
	Ip          string `json:"ip"`
	Port        int    `json:"port"`
	Protocol    string `json:"protocol"`
	Vulntype    int    `json:"vulntype"`
	User        string `json:"user"`
	Pass        string `json:"pass"`
	TaskId      int    `json:"task_id"`
	TaskTime    string `json:"task_time"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

type Service struct {
	Ip       string
	Port     int
	Protocol string
	UserName string
	PassWord string
}

type ScanResult struct {
	Service Service
	Result  bool
}

type IpAddr struct {
	Ip       string
	Port     int
	Protocol string
}

//创建任务
func AddPortBruteRes(data map[string]interface{}) {
	nowTime := time.Now().Format("20060102150405")
	portBruteRes := Portbruteres{
		Ip:       data["ip"].(string),
		Port:     data["port"].(int),
		User:     data["user"].(string),
		Protocol: data["protocol"].(string),
		Pass:     data["pass"].(string),
		TaskId:   data["taskid"].(int),
		//Vulntype:        data["vulntype"].(int),
		//TaskId:        data["task_id"].(string),
		TaskTime:    data["task_time"].(string),
		CreatedTime: nowTime,
		UpdatedTime: nowTime,
	}
	db.Create(&portBruteRes)
}

func GetPortBruteRes(pageNum int, pageSize int, maps interface{}) (portbruteres []Portbruteres) {
	db.Where(maps).Order("created_time").Find(&portbruteres)
	return
}

func GetPortBruteResTotal(maps interface{}) (count int) {
	db.Model(&Portbruteres{}).Where(maps).Count(&count)
	return
}

func GetPortBruteResTotalGroupBy() (portbruteres []Portbruteres) {
	var protocol string
	var num int

	rows, _ := db.Model(&Portbruteres{}).Select("protocol as protocol, COUNT(*) as num").Group("protocol").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&protocol, &num)
		portbruteresTmp := Portbruteres{
			Protocol: protocol,
			Vulntype: num,
		}
		portbruteres = append(portbruteres, portbruteresTmp)
	}

	return
}

func GetBruteResLastUpdate() (portbruteres []Portbruteres) {
	db.Limit(10).Order("updated_time desc").Find(&portbruteres)
	return
}

func GetPortBruteResById(task_id int) (portbruteres []Portbruteres) {
	db.Order("created_time").Where("task_id = ? ", task_id).Find(&portbruteres)
	return
}

func DeletePortBruteResById(taskId int) {
	db.Where("task_id = " + strconv.Itoa(taskId)).Delete(&Portbruteres{})
}
