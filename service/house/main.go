package main

import (
	"github.com/cloudwego/kitex/server"
	"ihome/service/house/conf"
	kitex_gen "ihome/service/house/kitex_gen/houseservice"
	"ihome/service/house/model"
	"ihome/service/utils"
	"net"
)

func main() {
	//初始化redis
	model.InitRedis()
	utils.NewLog().Info("init redis..", model.Client)
	//初始化mysql
	model.InitDb()
	utils.NewLog().Info("init mysql..", model.MysqlConn)
	svr := kitex_gen.NewServer(new(HouseServiceImpl),
		server.WithServiceAddr(&net.TCPAddr{Port: conf.ServerPort, IP: net.ParseIP(conf.ServerIp)}))
	err := svr.Run()
	if err != nil {
		utils.NewLog().Error("service start error", err)
	}
}
