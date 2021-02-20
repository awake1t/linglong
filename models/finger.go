package models

import (
	"strconv"
	"time"
)

type Finger struct {
	*Model
	Name        string `json:"name"`
	Description    string `json:"description"`
	Finger         string `json:"finger"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}



func GetFinger(pageNum int, pageSize int, maps interface{}) (finger []Finger) {
	dbTmp := db
	querys := maps.(map[string]interface{})
	if querys["name"] != nil {
		dbTmp = dbTmp.Where("name LIKE ?", "%"+querys["name"].(string)+"%")
	}
	if querys["finger"] != nil {
		dbTmp = dbTmp.Where("finger LIKE ?", "%"+querys["finger"].(string)+"%")
	}
	dbTmp.Offset(pageNum).Limit(pageSize).Order("updated_time  desc").Find(&finger)
	return
}



func GetFingerTotal(maps interface{}) (count int) {
	dbTmp := db
	querys := maps.(map[string]interface{})
	if querys["name"] != nil {
		dbTmp = dbTmp.Where("name LIKE ?", "%"+querys["name"].(string)+"%")
	}
	dbTmp.Model(&Finger{}).Where(maps).Count(&count)
	return
}


func GetAllFinger() (finger []Finger) {
	db.Find(&finger)
	return
}



func GetAllFingerId(id int) (finger []Finger) {
	dbTmp := db
	dbTmp.Where("id = ?", id).Find(&finger)
	return
}

func AddFinger(data map[string]interface{}) {
	nowTime := time.Now().Format("20060102150405")
	Finger := Finger{
		Name:          data["name"].(string),
		Description:          data["description"].(string),
		Finger:          data["finger"].(string),
		CreatedTime: nowTime,
		UpdatedTime: nowTime,
	}
	db.Create(&Finger)
}



func ExistFinger(ip, port string) (bool, int) {
	var Finger Finger
	db.Select("id").Where("ip = ? and port = ? ", ip, port).First(&Finger)
	//如果返回的id>0，也就是数据库里存在过了数据
	if Finger.ID > 0 {
		return true, Finger.ID
	}

	return false, Finger.ID
}

//通过id，更新ip列表
func EditFinger(id int, data interface{}) bool {
	db.Model(&Finger{}).Where("id = ?", id).Updates(data)
	return true
}


func DeleteFinger(id int) {
	db.Where("id = " + strconv.Itoa(id)).Delete(&Finger{})
}
