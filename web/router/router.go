package router

import (
	"github.com/gin-gonic/gin"
	"ihome/web/conf"
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
	router.Static("/static", "view")
	router.MaxMultipartMemory = conf.UploadFileMaxSize
	router.GET("/", controller.Index)
	v1 := router.Group("/api/v1.0")
	{

		v1.GET("/imagecode/:uuid", controller.GetImageCd)
		v1.GET("/smscode/:phone", controller.GetSMSCd)
		v1.POST("/users", controller.Register)
		v1.GET("/areas", controller.GetAreas) //获取房子地域信息
		v1.GET("/test", controller.Test)

		v1.POST("/sessions", controller.PostLogin) //登录
		v1.GET("/session", controller.GetSession)  //获取用户信息

		v1.Use(controller.SessionAuth)
		//下面的方法都得sessionAuth
		v1.GET("/user", controller.GetUserInfo) //获取用户信息
		v1.PUT("/user/name", controller.UpdateUserInfo)
		v1.DELETE("/session", controller.DeleteSession)
		v1.POST("/user/avatar", controller.PostAvatar)
		v1.GET("/user/auth", controller.GetUserInfo)
		v1.POST("/user/auth", controller.UpdateUserInfo)
		//house
		v1.GET("/house/index", controller.HomePageIndex) //获取首页轮播
		v1.GET("/user/houses", controller.GetUserHouses) //获取用户发布的所有房子信息
		v1.POST("/houses", controller.PubHouses)         //用户发布房子信息
		v1.GET("/houses", controller.SearchHouse)        //搜索房源
		v1.POST("/houses/:id/images", controller.UploadHouseImg)
		v1.GET("/houses/:id", controller.GetHouseDetail)

		//订单
		v1.POST("/orders", controller.PostOrders)
	}

	return router
}
