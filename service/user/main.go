package main

import (
	"github.com/cloudwego/kitex/server"
	"ihome/service/user/conf"
	kitex_gen "ihome/service/user/kitex_gen/userservice"
	"ihome/service/user/model"
	"ihome/service/user/utils"
	"net"
)

func main() {
	model.InitRedis()
	utils.NewLog().Info("init redis..", model.Client)
	svr := kitex_gen.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(&net.TCPAddr{Port: conf.ServerPort, IP: net.ParseIP(conf.ServerIp)}))
	err := svr.Run()
	if err != nil {
		utils.NewLog().Error("service start error", err)
	}
}
