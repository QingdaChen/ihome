package router

import (
	"github.com/gin-gonic/gin"
	"ihome/web/controller"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	//使用session
	// 初始化容器.
	//映射静态资源
	//fmt.Println(os.Getwd())
	//Chdir()
	//fmt.Println("111:")
	//fmt.Println(os.Getwd())
	//windows
	//router.Static("/home", "./web/view")

	//linux
	//router.Static("/home/login", "view/login.html")
	router.Static("/home", "view")

	v1 := router.Group("/api/v1.0")
	{

		v1.GET("/imagecode/:uuid", controller.GetImageCd)
		v1.GET("/smscode/:phone", controller.GetSMSCd)
		v1.POST("/users", controller.Register)
		v1.GET("/areas", controller.GetAreas) //获取房子地域信息

		v1.POST("/sessions", controller.PostLogin) //登录
		v1.GET("/session", controller.GetSession)  //获取用户信息

		//下面的方法都得sessionAuth
		v1.Use(controller.SessionAuth(router))
		v1.GET("/user", controller.GetUserInfo) //获取用户信息
		v1.PUT("/user/name", controller.UpdateUserInfo)
		v1.DELETE("/session", controller.DeleteSession)
		v1.POST("/user/avatar", controller.PostAvatar)

	}
	//router.Use(SessionAuthorize(router))
	//以下方法均使用session鉴权

	return router
}
