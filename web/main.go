package main

import (
	"ihome/web/controller"
	"ihome/web/model"
	"ihome/web/utils"

	"github.com/gin-gonic/gin"
)

func main() {
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
	}
	//初始化redis
	model.InitDb()
	//初始化缓存
	utils.InitCache()
	//TODO wire 依赖注入 service
	router.Run(":8088")
}
