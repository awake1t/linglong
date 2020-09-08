package v1

import (
	"github.com/gin-gonic/gin"
	"linglong/models"
	"linglong/pkg/e"
	"linglong/pkg/utils"
	"net/http"
)

//任务列表
func GetLogList(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	code := e.SUCCESS

	data["lists"] = models.GetLog(utils.GetPage(c), 10, maps)
	data["total"] = models.GetLogTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
