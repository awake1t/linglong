package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"linglong/models"
	"linglong/pkg/e"
	"linglong/pkg/utils"
	"net/http"
	"strings"
)

func GetFinger(c *gin.Context) {
	name := c.Query("name")
	finger := c.Query("finger")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	if finger != "" {
		maps["finger"] = finger
	}

	code := e.SUCCESS

	data["lists"] = models.GetFinger(utils.GetPage(c), 10, maps)
	data["total"] = models.GetFingerTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddFinger(c *gin.Context) {

	name := c.PostForm("name")
	description := c.PostForm("description")
	finger := c.PostForm("finger")

	code := e.INVALID_PARAMS

	if strings.HasSuffix(finger, ",") {
		code = e.INVALID_FINGER
	} else {

		valid := validation.Validation{}

		// 输入长度限制
		valid.Required(name, "name").Message("名称不能为空")

		if !valid.HasErrors() {

			go func() {
				data := make(map[string]interface{})
				data["name"] = name
				data["description"] = description
				data["finger"] = finger
				models.AddFinger(data)
			}()
			code = e.SUCCESS

		}

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})


}

// 新增指纹时测试指纹是否正确
func TestFinger(c *gin.Context) {

	testurl := c.PostForm("testurl")
	finger := c.PostForm("finger")

	valid := validation.Validation{}

	// 输入长度限制
	valid.Required(finger, "finger").Message("finger不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if CheckFingerSingle(testurl, finger) {
			code = e.SUCCESS
		} else {
			code = 402
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// 扫描当前指纹在数据库中的数量
func ScanFinger(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	code := e.SUCCESS
	CheckFingerOne(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

func DeleteFinger(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {

		//删除任务结果记录
		models.DeleteFinger(id)
		code = e.SUCCESS

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}
