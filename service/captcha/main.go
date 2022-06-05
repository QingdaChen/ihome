package main

import (
	"github.com/cloudwego/kitex/server"
	"ihome/service/captcha/conf"
	kitex_gen "ihome/service/captcha/kitex_gen/captchaservice"
	"ihome/service/captcha/model"
	"ihome/service/captcha/utils"
	"net"
)

func main() {
	model.InitRedis()
	utils.NewLog().Info("main pool", model.Client)
	svr := kitex_gen.NewServer(new(CaptchaServiceImpl),
		server.WithServiceAddr(&net.TCPAddr{Port: conf.ServerPort, IP: net.ParseIP(conf.ServerIp)}))
	err := svr.Run()
	if err != nil {
		utils.NewLog().Error("service start error...", err)
	}

}
