package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"linglong/models"
	"linglong/pkg/brute"
	"linglong/pkg/common"
	"strconv"
	"time"
)

//新建端口爆破
//异步执行
/*
参数： 1. 资产来源 (数据库资产/输入资产) 2.资产内容 3. 密码类型(通用字典/数据库选择/临时输入) 4。用户名字典  5 密码字典 6.线程 int
7. 定时任务id	(比如ssh弱口令）		定时任务周期(比如 上周结果存入一个数据库，这周结果存入另一个库
返回: 1. 任务状态（0 已创建；1 队列中；2 扫描中；3 已完成；4 任务失败；5 任务停止'）
*/
func NewPortBrute(c *gin.Context) {

	//资产来源
	source := com.StrTo(c.PostForm("source")).MustInt()
	//资产内容
	iplist := c.PostForm("iplist")
	dict := com.StrTo(c.PostForm("dict")).MustInt()
	//user := c.PostForm("user")
	//pass := c.PostForm("pass")
	thread := com.StrTo(c.PostForm("thread")).MustInt()
	taskid := com.StrTo(c.PostForm("taskid")).MustInt()


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

