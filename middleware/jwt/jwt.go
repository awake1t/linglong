package jwt

import (
	"github.com/gin-gonic/gin"
	"linglong/pkg/e"
	"linglong/pkg/utils"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.GetHeader("Authorization")

		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := utils.ParseToken(token)
			if claims == nil{
				c.JSON(http.StatusUnauthorized, gin.H{
					"code" : 401,
					"msg" : "cookie失效，请点击右上角退出重新登陆",
					"data" : data,
				})
				c.Abort()
				return
			}
			if err != nil {
				code = e.ERROR
			} else if time.Now().Unix() > claims.ExpiresAt {
				code =e.ERROR
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
