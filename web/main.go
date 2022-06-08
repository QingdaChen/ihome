package main

import (
	"ihome/web/conf"
	"ihome/web/model"
	"ihome/web/router"
	"ihome/web/utils"
	"strconv"
)

func main() {

	//初始化redis
	model.InitDb()
	//初始化缓存
	utils.InitCache()
	//TODO wire 依赖注入 service
	r := router.InitRouters()
	r.Run(":" + strconv.Itoa(conf.WebPort))
}
