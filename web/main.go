package main

import (
	"github.com/gin-gonic/gin"
	"ihome/web/controller"
	"ihome/web/model"
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
	//开启监听
	//router.GET("/", func(context *gin.Context) {
	//	context.Writer.WriteString("hhhh")
	//})
	v1 := router.Group("/api/v1.0")
	{
		v1.GET("/session", controller.GetSession)
		v1.GET("/imagecode/:uuid", controller.GetImageCd)
		v1.GET("/smscode/:phone", controller.GetSMSCd)
	}
	//初始化redis
	model.InitDb()
	router.Run(":8088")
}
