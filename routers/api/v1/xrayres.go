package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"linglong/models"
	"linglong/pkg/e"
	"linglong/pkg/utils"
	"net/http"
)

func GetXrayres(c *gin.Context) {
	url := c.Query("url")
	poc := c.Query("poc")
	snapshot := c.Query("snapshot")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if url != "" {
		maps["url"] = url
	}
	if poc != "" {
		maps["poc"] = poc
	}
	if snapshot != "" {
		maps["snapshot"] = snapshot
	}

	code := e.SUCCESS

	data["lists"] = models.GetXrayres(utils.GetPage(c), 10, maps)
	data["total"] = models.GetXrayresTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func DeleteXrayres(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		models.DeleteXrayres(id)
		code = e.SUCCESS

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}
