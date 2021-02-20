package models

import (
	"strconv"
	"time"
)

type Xrayres struct {
	*Model
	Url         string `json:"url"`
	Poc         string `json:"poc"`
	Snapshot    string `json:"snapshot"`
	Hash        string `json:"hash"`
	CreatedTime string `json:"created_time"`
}

func GetXrayres(pageNum int, pageSize int, maps interface{}) (Xrayres []Xrayres) {
	dbTmp := db
	querys := maps.(map[string]interface{})

	if querys["url"] != nil {
		dbTmp = dbTmp.Where("url LIKE ?", "%"+querys["url"].(string)+"%")
	}

	if querys["poc"] != nil {
		dbTmp = dbTmp.Where("poc LIKE ?", "%"+querys["poc"].(string)+"%")
	}

	if querys["snapshot"] != nil {
		dbTmp = dbTmp.Where("snapshot LIKE ?", "%"+querys["snapshot"].(string)+"%")
	}

	dbTmp.Offset(pageNum).Limit(pageSize).Order("created_time  desc").Find(&Xrayres)
	return
}

func GetXrayresTotal(maps interface{}) (count int) {
	dbTmp := db
	querys := maps.(map[string]interface{})

	if querys["url"] != nil {
		dbTmp = dbTmp.Where("url LIKE ?", "%"+querys["url"].(string)+"%")
	}

	if querys["poc"] != nil {
		dbTmp = dbTmp.Where("poc LIKE ?", "%"+querys["poc"].(string)+"%")
	}

	if querys["snapshot"] != nil {
		dbTmp = dbTmp.Where("snapshot LIKE ?", "%"+querys["snapshot"].(string)+"%")
	}
	dbTmp.Model(&Xrayres{}).Count(&count)
	return
}

//创建任务，返回任务id
func AddXrayres(data map[string]interface{}) {
	nowTime := time.Now().Format("20060102150405")
	Xrayres := Xrayres{
		Url:         data["url"].(string),
		Poc:         data["poc"].(string),
		Snapshot:    data["snapshot"].(string),
		Hash:        data["hash"].(string),
		CreatedTime: nowTime,
	}
	db.Create(&Xrayres)
}


func DeleteXrayres(id int) {
	db.Where("id = " + strconv.Itoa(id)).Delete(&Xrayres{})
}
