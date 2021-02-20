package routers

import (
	"github.com/gin-gonic/gin"
	"linglong/global"
	"linglong/middleware/jwt"
	"linglong/routers/api"
	"linglong/routers/api/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.Use(Cors())

	gin.SetMode(global.ServerSetting.RunMode)

	r.GET("/api/v1/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")

	//r.LoadHTMLGlob("dist/*.html")              // 添加入口index.html
	//r.LoadHTMLFiles("./css/*")                   // 添加资源路径
	//r.Static("/css", "./dist/css")             // 添加资源路径
	//r.LoadHTMLFiles("fonts/*")                 // 添加资源路径
	//r.Static("/fonts", "./dist/fonts")         // 添加资源路径
	//r.LoadHTMLFiles("img/*")                   // 添加资源路径
	//r.Static("/img", "./dist/img")             // 添加资源路径
	//r.LoadHTMLFiles("js/*")                    // 添加资源路径
	//r.Static("/js", "./dist/js")               // 添加资源路径
	//r.StaticFile("/", "dist/index.html") //前端接口

	//新建端口爆破 - 无认证
	//apiv1.POST("/nweport", v1.NewPortBrute)
	apiv1.GET("/downtasklog/:id", v1.DownTaskLog)


	apiv1.Use(jwt.JWT())

	{
		// masscan资产任务列表
		apiv1.GET("/masstasks", v1.GetIplist)
		// ip资产列表搜索
		apiv1.GET("/masstask", v1.GetIplist)
		//新建
		apiv1.POST("/masstask", v1.AddIplist)

		//定时任务列表
		apiv1.GET("/crons",v1.GetTask)
		apiv1.POST("/addcron",v1.AddTask)
		apiv1.POST("/startcron",v1.StartTask)
		apiv1.DELETE("/delcron/:id",v1.DeleteTask)

		// 获取日志列表
		apiv1.GET("/log", v1.GetLogList)
		// 获取指定任务taskid结果
		apiv1.GET("/masstasklog/:id", v1.GetTaskLog)

		// 传入任务执行时间，返回任务结果(爆破结果数据库）
		apiv1.GET("/masstasktime", v1.GetTaskTime)

		// 传入任务执行时间，返回任务状态(任务状态数据库)
		apiv1.GET("/masstaskstatus", v1.GetTaskStatus)



		//平台设置
		apiv1.GET("/setting", v1.GetSetting)
		//修改设置
		apiv1.PUT("/setting", v1.EditSetting)
		//修改密码
		apiv1.PUT("/modpass", v1.EditPass)

		// 管理后台列表
		apiv1.GET("/webloginlist", v1.GetWebloginlist)
		// 根据title搜索管理后台
		apiv1.GET("/webloginlistsearch", v1.GetWebloginlist)

		//数据统计页面接口
		apiv1.GET("/dashboard", v1.Getdashboard)


		apiv1.GET("/finger", v1.GetFinger)
		apiv1.POST("/addfinger", v1.AddFinger)
		// 测试指纹
		apiv1.POST("/testfinger", v1.TestFinger)
		apiv1.GET("/scanfinger/:id", v1.ScanFinger)
		apiv1.DELETE("/delfinger/:id",v1.DeleteFinger)


		apiv1.GET("/gerXrayRes", v1.GetXrayres)
		apiv1.DELETE("/delxrayres/:id",v1.DeleteXrayres)

	}
	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization,Token,X-TOKEN ")
		c.Header("Access-Control-Allow-Methods", "POST, GET,PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
