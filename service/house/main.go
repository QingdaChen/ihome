package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/server"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"ihome/service/house/conf"
	kitex_gen "ihome/service/house/kitex_gen/houseservice"
	"ihome/service/utils"
	"net"
)

func main() {
	////初始化redis
	//model.InitRedis()
	//utils.NewLog().Info("init redis..", model.Client)
	////初始化mysql
	//model.InitDb()
	//utils.NewLog().Info("init mysql..", model.MysqlConn)
	svr := kitex_gen.NewServer(new(HouseServiceImpl),
		server.WithServiceAddr(&net.TCPAddr{Port: conf.ServerPort, IP: net.ParseIP(conf.ServerIp)}),
		server.WithLimit(&limit.Option{MaxConnections: conf.ServerMaxConnections, MaxQPS: conf.ServerMaxQPS}), // limit
		server.WithMuxTransport(),                       // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite())) // tracer
	err := svr.Run()
	if err != nil {
		utils.NewLog().Error("service start error", err)
	}
}
