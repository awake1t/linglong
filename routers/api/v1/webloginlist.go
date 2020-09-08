package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"linglong/models"
	"linglong/pkg/e"
	"linglong/pkg/utils"
	"log"
	"net/http"
)

func GetWebloginlist(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	code := e.SUCCESS

	data["lists"] = models.GetWebloginlist(utils.GetPage(c), 10, maps,"")
	data["total"] = models.GetWebloginlistTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//查询
func GetWebloginlistSearch(c *gin.Context) {
	code := e.INVALID_PARAMS
	valid := validation.Validation{}
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	title := c.Query("title")

	if ! valid.HasErrors() {
		data["lists"] = models.GetWebloginlist(utils.GetPage(c), 10, maps,title)
		data["total"] = models.GetWebloginlistTotal(maps)
		code = e.SUCCESS
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
