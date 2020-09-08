package models

import (
	"time"
)

type Log struct {
	Id          int    `gorm:"primary_key" json:"id"`
	TaskId      int    `json:"task_id"`
	TaskName    string `json:"task_name"`
	TaskType    string `json:"task_type"`
	AllNum      int    `json:"all_num"`
	SuccesNum   int    `json:"succes_num"`
	RunTime     string `json:"run_time"`
	Status      int    `json:"status"`
	Error       string `json:"error"`
	CreatedTime string `json:"created_time"`
}


func GetLog(pageNum int, pageSize int, maps interface{}) (log []Log) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Order("created_time desc").Find(&log)
	return
}

func GetLogTotal(maps interface{}) (count int) {
	db.Model(&Log{}).Where(maps).Count(&count)

	return
}

func AddLog(data map[string]interface{}) {
	log := Log{
		TaskId:      data["taskid"].(int),
		TaskName:    data["task_name"].(string),
		TaskType:    data["task_type"].(string),
		AllNum:      data["all_num"].(int),
		SuccesNum:   data["succes_num"].(int),
		RunTime:     data["run_time"].(string),
		Error:       data["error"].(string),
		Status:      data["status"].(int),
		CreatedTime: time.Now().Format("20060102150405"),
	}
	db.Create(&log)
}

func GetLogUpdate() (log []Log) {
	db.Select("created_time").Order("created_time desc").Limit(10).Find(&log)
	return
}

