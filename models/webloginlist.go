package models

import "time"

type Webloginlist struct {
	*Model
	Ip          string `json:"ip"`
	Port        string `json:"port"`
	Protocol    string `json:"protocol"`
	Url         string `json:"url"`
	Title       string `json:"title"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

func GetWebloginlist(pageNum int, pageSize int, maps interface{},title string) (webloginlist []Webloginlist) {
	db.Where(maps).Where("title LIKE ?", "%"+title+"%").Offset(pageNum).Limit(pageSize).Find(&webloginlist)
	return
}



func GetWebloginlistTotal(maps interface{}) (count int) {
	db.Model(&Webloginlist{}).Where(maps).Count(&count)
	return
}

func AddWebloginlist(data map[string]interface{}) {
	nowTime := time.Now().Format("20060102150405")
	webloginlist := Webloginlist{
		Ip:          data["ip"].(string),
		Port:        data["port"].(string),
		Protocol:    data["protocol"].(string),
		Url     :    data["url"].(string),
		Title    :   data["title"].(string),
		CreatedTime: nowTime,
		UpdatedTime: nowTime,
	}
	db.Create(&webloginlist)
}



func ExistWebloginlist(ip, port string) (bool, int) {
	var webloginlist Webloginlist
	db.Select("id").Where("ip = ? and port = ? ", ip, port).First(&webloginlist)
	//如果返回的id>0，也就是数据库里存在过了数据
	if webloginlist.ID > 0 {
		return true, webloginlist.ID
	}

	return false, webloginlist.ID
}

//通过id，更新ip列表
func EditWebloginlist(id int, data interface{}) bool {
	db.Model(&Webloginlist{}).Where("id = ?", id).Updates(data)
	return true
}

//获取最近更新的title
func GetWebloginLastUpdate() (webloginlist []Webloginlist) {
	db.Limit(10).Find(&webloginlist)
	return
}
