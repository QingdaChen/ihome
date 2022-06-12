package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/server"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"ihome/service/captcha/conf"
	kitex_gen "ihome/service/captcha/kitex_gen/captchaservice"
	"ihome/service/utils"
	"net"
)

func main() {
	//model.InitRedis()
	//utils.NewLog().Info("main pool", model.Client)
	svr := kitex_gen.NewServer(new(CaptchaServiceImpl),
		server.WithServiceAddr(&net.TCPAddr{Port: conf.ServerPort, IP: net.ParseIP(conf.ServerIp)}),
		server.WithLimit(&limit.Option{MaxConnections: conf.ServerMaxConnections, MaxQPS: conf.ServerMaxQPS}), // limit
		server.WithMuxTransport(),                       // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite())) // tracer
	err := svr.Run()
	if err != nil {
		utils.NewLog().Error("service start error...", err)
	}

}
