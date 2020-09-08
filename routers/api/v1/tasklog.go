package v1

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"linglong/models"
	"linglong/pkg/e"
	"linglong/pkg/utils"
	"log"
	"net/http"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

func GetTaskLog(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		maps["task_id"] = id
		code = e.SUCCESS
		data["lists"] = models.GetTaskLog(maps)
		data["total"] = models.GetTaskLogTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

//传入任务执行时间，返回任务状态(任务状态数据库)
func GetTaskStatus(c *gin.Context) {
	taskid := c.Query("taskid")
	tasktime := c.Query("tasktime")

	data := make(map[string]interface{})

	if tasktime == "0" {
		data["lists"] = models.GetTaskLogLastStatusById(taskid)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  e.GetMsg(200),
			"data": data,
		})
	} else {
		data["lists"] = models.GetTaskLogLastStatuByTime(tasktime)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  e.GetMsg(200),
			"data": data,
		})

	}
}

func GetTaskTime(c *gin.Context) {
	taskid := c.Query("taskid")
	tasktime := c.Query("tasktime")
	fmt.Println("taskid :", taskid)
	fmt.Println("tasktime :", tasktime)

	if tasktime == "0" {
		tmpTaskTime := models.GetTaskLogLast(taskid)
		fmt.Println(tmpTaskTime[0].CreatedTime)
		tasktime = tmpTaskTime[0].CreatedTime
	}

	valid := validation.Validation{}

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		maps["task_time"] = tasktime
		maps["task_id"] = taskid
		code = e.SUCCESS
		data["lists"] = models.GetPortBruteRes(utils.GetPage(c), 10, maps)
		data["total"] = models.GetPortBruteResTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

//下载任务结果
func DownTaskLog(c *gin.Context) {
	fmt.Println("DownTaskLog")
	id := com.StrTo(c.Param("id")).MustInt()
	fmt.Println("id:",id)

	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "ip")
	f.SetCellValue("Sheet1", "B1", "port")
	f.SetCellValue("Sheet1", "C1", "protocol")
	f.SetCellValue("Sheet1", "D1", "账号")
	f.SetCellValue("Sheet1", "E1", "密码")
	f.SetCellValue("Sheet1", "F1", "时间")

	res := models.GetPortBruteResById(id)
	fmt.Println("res:",res)
	for i,v := range res{
		fmt.Println("i:",i)
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), v.Ip)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), v.Port)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), v.Protocol)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), v.User)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+2), v.Pass)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+2), v.UpdatedTime)
	}
	if err := f.SaveAs("漏洞报告.xlsx"); err != nil {
		fmt.Println(err)
	}

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "漏洞报告.xlsx")) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./漏洞报告.xlsx")


}
