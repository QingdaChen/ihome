package router

import (
	"github.com/gin-gonic/gin"
	"ihome/web/controller"
)

func InitRouters() *gin.Engine {
	router := gin.Default()

	//映射静态资源
	//fmt.Println(os.Getwd())
	//Chdir()
	//fmt.Println("111:")
	//fmt.Println(os.Getwd())
	//windows
	//router.Static("/home", "./web/view")

	//linux
	router.Static("/home", "view")

	v1 := router.Group("/api/v1.0")
	{
		v1.GET("/session", controller.GetSession)
		v1.GET("/imagecode/:uuid", controller.GetImageCd)
		v1.GET("/smscode/:phone", controller.GetSMSCd)
		v1.POST("/users", controller.Register)
		v1.GET("/areas", controller.GetAreas)
	}
	return router
}