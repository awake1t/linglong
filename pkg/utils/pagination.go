package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"strconv"
)

func GetPage(c *gin.Context) int {
	result := 0
	pagesizetmp := c.Query("pagesize")
	pagesize, _ := strconv.Atoi(pagesizetmp)

	page, _ := com.StrTo(c.Query("pagenum")).Int()
	if page > 0 {
		result = (page - 1) * pagesize
	}

	return result
}
