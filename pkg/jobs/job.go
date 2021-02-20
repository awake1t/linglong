package jobs

import (
	"fmt"
	"linglong/models"
	"linglong/pkg/brute"
	"linglong/pkg/common"
	"strconv"
	"strings"
	"time"
)

type Job struct {
	id         int                                               // 任务ID
	logId      int64                                             // 日志记录ID
	name       string                                            // 任务名称
	task       *models.Task                                      // 任务对象
	runFunc    func(time.Duration) (string, string, error, bool) // 执行函数
	status     int                                               // 任务状态，大于0表示正在执行中
	Concurrent bool                                              // 同一个任务是否允许并行执行
}

func NewJobFromTask(task *models.Task) (*Job, error) {
	job := NewCommandJob(task)
	job.task = task
	job.Concurrent = task.Concurrent == 1
	return job, nil
}

func NewCommandJob(task *models.Task) *Job {
	job := &Job{
		id:   task.Id,
		name: task.TaskName,
	}
	job.runFunc = func(timeout time.Duration) (string, string, error, bool) {
		arges := strings.Split(task.Arge,",")
		thread ,_ := strconv.Atoi(arges[1])
		NewPortBrute(arges[0],1,1,thread,task.Id)
		return "", "", nil,true
	}
	return job
}

func (j *Job) Run() {
	if !j.Concurrent && j.status > 0 {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("执行任务失败", err)
		}
	}()

	if workPool != nil {
		workPool <- true
		defer func() {
			<-workPool
		}()
	}
	j.status++
	defer func() {
		j.status--
	}()

	timeout := time.Duration(time.Hour * 24)
	if j.task.Timeout > 0 {
		timeout = time.Second * time.Duration(j.task.Timeout)
	}

	_, _, err, _ := j.runFunc(timeout)

	if err !=nil{
		fmt.Println("eccadcjads",err)
	}


	// 如果是之执行一次的任务，修改任务状态为0，暂停任务
	if j.task.TaskCycle == "now"{
		data := make(map[string]interface{})
		data["status"] = 0
		models.EditTask(j.task.Id,data)
	}



}

func NewPortBrute(iplist string, source,dict,thread,taskid int) {

	ipList := make([]models.IpAddr, 0)
	if source == 1 {
		fmt.Println("从数据库选择资产 port:3306", iplist)
		var port int
		var protocol string
		if iplist == "mysql" {
			port = 3306
			protocol = "MYSQL"
		} else if iplist == "ssh" {
			port = 22
			protocol = "SSH"
		}else if iplist == "ftp" {
			port = 21
			protocol = "FTP"
		}else if iplist == "smb" {
			port = 445
			protocol = "SMB"
		}else if iplist == "mssql" {
			port = 1433
			protocol = "MSSQL"
		}else if iplist == "postgresql" {
			port = 5432
			protocol = "POSTGRESQL"
		}else if iplist == "mongodb" {
			port = 27017
			protocol = "MONGODB"
		}else if iplist == "redis" {
			port = 6379
			protocol = "REDIS"
		}
		// 从数据库取结果
		datalists := models.GetIplistBrute(port, protocol)
		for _, target := range datalists {
			tmpPort, _ := strconv.Atoi(target.Port)
			addr := models.IpAddr{Ip: target.Ip, Port: tmpPort, Protocol: protocol}
			ipList = append(ipList, addr)
		}

	} else if source == 2 {
		fmt.Println("临时输入资产")
	}

	if dict == 1 {
		//fmt.Println("默认通用字典")
	} else if dict == 2 {
		fmt.Println("灵活选择")
	} else if dict == 3 {
		fmt.Println("手动输入")
	}

	userDict, uErr := common.ReadUserDict("./configs/user.txt")
	passDict, pErr := common.ReadUserDict("./configs/pass.txt")
	taskruntime := time.Now().Format("20060102150405")

	if uErr == nil && pErr == nil {
		scanTasks := brute.GenerateTask(ipList, userDict, passDict)
		//爆破之前记录任务id到日志
		data := make(map[string]interface{})
		data["taskid"] = taskid
		data["created_time"] = taskruntime
		data["all_num"] = len(ipList)
		data["userdict"] = 0
		data["passdict"] = 0
		models.AddTaskLog(data)
		brute.RunTask(scanTasks, thread, taskid,taskruntime)
	} else {
		fmt.Println("Read File Err!")
	}

	//根据任务id 取出爆破成功数量
	maps := make(map[string]interface{})
	maps["task_id"] = taskid
	succesnum := models.GetPortBruteResTotal(maps)
	data := make(map[string]interface{})
	data["vuln_num"] = succesnum
	models.EditTask(taskid, data)

	mapsTaskLog := make(map[string]interface{})
	mapsTaskLog["task_id"] = taskid
	succesnumTaskLog := models.GetPortBruteResTotal(mapsTaskLog)

	dataTaskLog := make(map[string]interface{})
	dataTaskLog["vuln_num"] = succesnumTaskLog
	models.EditTaskLogTaskTime(taskid, taskruntime,dataTaskLog)

}

