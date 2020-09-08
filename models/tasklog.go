package models

type TaskLog struct {
	Id          int    `gorm:"primary_key" json:"id"`
	TaskId      int    `json:"task_id"`
	Output      string `json:"output"`
	Error       string `json:"error"`
	Status      int    `json:"status"`
	AllNum      int    `json:"all_num"`
	SuccesNum   int    `json:"succes_num"`
	Userdict    int    `json:"userdict"`
	Passdict    int    `json:"passdict"`
	RunTime     string    `json:"run_time"`
	ProcessTime int    `json:"process_time"`
	CreatedTime string `json:"created_time"`
}


func GetTaskLogLast(taskid string) (taskLog []TaskLog) {
	db.Select("created_time").Where("task_id = ?", taskid).Order("created_time desc").First(&taskLog)
	return
}


func GetTaskLog(maps interface{}) (taskLog []TaskLog) {
	db.Select("created_time").Where(maps).Order("created_time desc").Find(&taskLog)
	return
}

func GetTaskLogTotal(maps interface{}) (count int) {
	db.Model(&TaskLog{}).Where(maps).Count(&count)

	return
}

func EditTaskLogTaskId(id int, data interface{}) bool {
	db.Model(&TaskLog{}).Where("task_id = ?", id).Updates(data)
	return true
}

func EditTaskLogTaskTime(id int,createTime string, data interface{}) bool {
	db.Model(&Task{}).Where("id = ? AND created_time = ?", id,createTime).Updates(data)
	return true
}

func EditTaskLog(id int, data interface{}) bool {
	db.Model(&TaskLog{}).Where("id = ?", id).Updates(data)
	return true
}



func AddTaskLog(data map[string]interface{}) {
	task := TaskLog{
		TaskId:      data["taskid"].(int),
		CreatedTime: data["created_time"].(string),
		AllNum:      data["all_num"].(int),
		Userdict:    data["userdict"].(int),
		Passdict:    data["passdict"].(int),
		Status:      1,
		ProcessTime: 0,
	}
	db.Create(&task)
}


func GetTaskLogLastStatusById(taskid string) (taskLog []TaskLog) {
	db.Where("task_id = ?", taskid).Order("created_time desc").First(&taskLog)
	return
}


func GetTaskLogLastStatuByTime(tasktime string) (taskLog []TaskLog) {
	db.Where("created_time = ?", tasktime).Order("created_time desc").First(&taskLog)
	return
}
