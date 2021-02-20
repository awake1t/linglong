package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"linglong/models"
	"linglong/pkg/e"
	"linglong/pkg/utils"
	"net/http"
	"strconv"
	"time"
)

//任务列表
func GetSetting(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.SUCCESS

	data["lists"] = models.GetSetting(utils.GetPage(c), 10, maps)
	data["total"] = models.GetSettingTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//修改
func EditSetting(c *gin.Context) {

	login_word := c.PostForm("login_word")
	login_url := c.PostForm("login_url")
	masscan_ip := c.PostForm("masscan_ip")
	masscan_thred := com.StrTo(c.PostForm("masscan_thred")).MustInt()
	masscan_port := c.PostForm("masscan_port")
	masscan_white := c.PostForm("masscan_white")
	masscan_deltime := com.StrTo(c.PostForm("masscan_deltime")).MustInt()

	valid := validation.Validation{}

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		data := make(map[string]interface{})
		data["login_word"] = login_word
		data["login_url"] = login_url
		data["masscan_thred"] = masscan_thred
		data["masscan_deltime"] = masscan_deltime
		data["masscan_ip"] = masscan_ip
		data["masscan_port"] = masscan_port
		data["masscan_white"] = masscan_white
		models.EditSetting(data)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 修改密码
func EditPass(c *gin.Context) {

	username := c.PostForm("username")
	oldpass := c.PostForm("oldpass")
	newpass := c.PostForm("newpass")
	newpass2 := c.PostForm("newpass2")

	code := e.INVALID_DIFFPASS
	if newpass != newpass2{
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  "两次密码不一致",
			"data": make(map[string]string),
		})
		return
	}

	valid := validation.Validation{}

	code = e.INVALID_PASS
	if ! valid.HasErrors() {
		isExist := models.CheckAuth(username, oldpass)
		if isExist {
			data := make(map[string]interface{})
			data["password"] = newpass
			models.EditAuth(username,data)
			code = e.SUCCESS
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

type node struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Dynamic struct {
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

// 数据展示
func Getdashboard(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	data["iplist"] = models.GetIplistTotal(maps)
	data["ip"] = models.GetIplistTotalDistinct()
	data["weblogin"] = models.GetWebloginlistTotal(maps)
	data["vuln"] = models.GetPortBruteResTotal(maps)

	// 获取饼图的结果
	res := models.GetPortBruteResTotalGroupBy()
	var nodes []node
	if len(res) == 0{
		node := node{
			Name:  "testData",
			Value: 1,
		}
		nodes = append(nodes, node)
	}

	for _, v := range res {
		node := node{
			Name:  v.Protocol,
			Value: v.Vulntype,
		}
		nodes = append(nodes, node)
	}
	data["vulnratio"] = nodes

	// 获取柱状图结果 - 近7日新增
	timeLineX := make([]string, 0)
	timeLineY := make([]int, 0)

	nowTime := time.Unix(time.Now().Unix(), 0).Format("2006-01-02")
	nowTime1 := time.Now().Add(-time.Hour * 24).Format("2006-01-02")
	nowTime2 := time.Now().Add(-time.Hour * 48).Format("2006-01-02")
	nowTime3 := time.Now().Add(-time.Hour * 72).Format("2006-01-02")
	nowTime4 := time.Now().Add(-time.Hour * 96).Format("2006-01-02")
	nowTime5 := time.Now().Add(-time.Hour * 120).Format("2006-01-02")
	nowTime6 := time.Now().Add(-time.Hour * 144).Format("2006-01-02")

	time := models.GetIplistGroupByDay(nowTime)
	time1 := models.GetIplistGroupByDay(nowTime1)
	time2 := models.GetIplistGroupByDay(nowTime2)
	time3 := models.GetIplistGroupByDay(nowTime3)
	time4 := models.GetIplistGroupByDay(nowTime4)
	time5 := models.GetIplistGroupByDay(nowTime5)
	time6 := models.GetIplistGroupByDay(nowTime6)

	timeLineX = append(timeLineX, nowTime6)
	timeLineY = append(timeLineY,time6)
	timeLineX = append(timeLineX, nowTime5)
	timeLineY = append(timeLineY,time5)
	timeLineX = append(timeLineX, nowTime4)
	timeLineY = append(timeLineY,time4)
	timeLineX = append(timeLineX, nowTime3)
	timeLineY = append(timeLineY,time3)
	timeLineX = append(timeLineX, nowTime2)
	timeLineY = append(timeLineY,time2)
	timeLineX = append(timeLineX, nowTime1)
	timeLineY = append(timeLineY,time1)
	timeLineX = append(timeLineX, nowTime)
	timeLineY = append(timeLineY,time)


	data["timelinex"] = timeLineX
	data["timeliney"] = timeLineY


	// 获取柱状图结果 - 端口
	portLine := models.GetIplistGroupBy()
	portLineX := make([]string, 0)
	portLineY := make([]int, 0)
	for _, v := range portLine {
		portLineX = append(portLineX, v.Port)
		portLineY = append(portLineY, v.Portnum)
	}
	data["portlinex"] = portLineX
	data["portliney"] = portLineY

	// 获取柱状图结果 - 服务
	protocolLine := models.GetIplistGroupByProtocol()
	protocolLineX := make([]string, 0)
	protocolLineY := make([]int, 0)
	for _, v := range protocolLine {
		protocolLineX = append(protocolLineX, v.Protocol)
		protocolLineY = append(protocolLineY, v.Portnum)
	}
	data["protocolinex"] = protocolLineX
	data["protocoliney"] = protocolLineY


	// 获取动态
	dynamicst := models.GetBruteResLastUpdate()
	var dynamicsvue []Dynamic
	for _, v := range dynamicst {
		node := Dynamic{
			Content:   "新漏洞 "+v.Ip+":"+strconv.Itoa(v.Port),
			Timestamp: v.UpdatedTime,
		}
		dynamicsvue = append(dynamicsvue, node)
	}
	data["dynamics"] = dynamicsvue


	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  e.GetMsg(200),
		"data": data,
	})

}
