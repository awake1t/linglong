package v1

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	libcron "github.com/robfig/cron/v3"
	"github.com/unknwon/com"
	"linglong/models"
	"linglong/pkg/e"
	"linglong/pkg/jobs"
	"linglong/pkg/utils"
	"net/http"
	"strconv"
	"time"
)

func GetTask(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	code := e.SUCCESS

	data["lists"] = models.GetTask(utils.GetPage(c), 10, maps)
	data["total"] = models.GetTaskTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 新增定时任务
func AddTask(c *gin.Context) {

	var arge string

	taskname := c.PostForm("taskname")
	description := c.PostForm("description")

	//配合前端灵活输入 最终得到corntab命令， 只有cronspec参数入库
	cronspec := c.PostForm("cronspec")
	day := c.PostForm("day")
	hour := c.PostForm("hour")
	cronspecmd := c.PostForm("cronspecmd")

	//爆破类型
	brute := c.PostForm("brute")
	source := c.PostForm("source")
	sourcecontent := c.PostForm("sourcecontent")

	thread := c.PostForm("thread")

	command := "http://127.0.0.1:18000/api/v1/nweport"
	valid := validation.Validation{}

	var taskcycle string
	if cronspec == "now" {
		nowMin := time.Now().Format("1504")
		min, _ := strconv.Atoi(nowMin[2:])
		min = min + 1
		cronspec = strconv.Itoa(min) + " " + nowMin[:2] + " * * *"
		fmt.Println("执行一次:",cronspec)
		taskcycle = "now"
	} else if cronspec == "day" {
		cronspec = "0 " + hour + " * * *"
		taskcycle = "每天"+hour+"点"
		fmt.Println("cronspec:", cronspec)
	} else if cronspec == "week" {
		cronspec = "0 " + hour + " * * " + day + ""
		fmt.Println("cronspec:", cronspec)
		taskcycle = "每周"+day+"的"+hour+"点"
	} else if cronspec == "cmd" {
		cronspec = cronspecmd
	} else {
	}

	if source == "1" {
		fmt.Println("brute:", brute)
		sourcecontent = brute
	} else if source == "2" {
	}
	arge = "source=" + source + "&iplist=" + sourcecontent + "&dict=1&thread=" + thread
	// 输入长度限制
	valid.Required(taskname, "taskname").Message("名称不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if _, err := libcron.ParseStandard(cronspec); err != nil {
			code = e.ERROR_CRON_SPEC
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": make(map[string]string),
			})
			return
		}

		go func() {

			data := make(map[string]interface{})
			data["taskname"] = taskname
			data["description"] = description
			data["cronspec"] = cronspec
			data["command"] = command
			data["arge"] = arge
			data["tasktype"] = brute
			data["taskcycle"] = taskcycle
			taskId := models.AddTask(data)

			//数据添加数据库后，根据任务id,立即执行任务
			task := models.GetTaskById(taskId)
			job, err := jobs.NewJobFromTask(task)
			if err != nil {
				fmt.Printf(" NewJobFromTask err ", err)
				return
			}

			jobs.AddJob(task.CronSpec, job)
		}()

		code = e.SUCCESS

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// 启动定时任务
func StartTask(c *gin.Context) {

	id := com.StrTo(c.PostForm("id")).MustInt()

	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {

		task := models.GetTaskById(id)
		job, err := jobs.NewJobFromTask(task)
		if err != nil {
			fmt.Printf(" NewJobFromTask err ", err)
			return
		}
		fmt.Printf(" jobssss %T \n %s \n", job, job)

		//第一次直接执行
		job.Run()

		//后期的定时循环
		jobs.AddJob(task.CronSpec, job)

		data := make(map[string]interface{})
		data["status"] = 1
		models.EditTask(id, data)

		code = e.SUCCESS

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// 删除任务
func DeleteTask(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	fmt.Println("idddsfsadassad is", id)

	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {

		models.DeleteTaskById(id)
		jobs.RemoveJob(id)

		//删除任务结果记录
		models.DeletePortBruteResById(id)

		code = e.SUCCESS

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}


