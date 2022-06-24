package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/server"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"ihome/service/order/conf"
	kitex_gen "ihome/service/order/kitex_gen/orderservice"
	"log"
	"net"
)

func main() {
	svr := kitex_gen.NewServer(new(OrderServiceImpl),
		server.WithServiceAddr(&net.TCPAddr{Port: conf.ServerPort, IP: net.ParseIP(conf.ServerIp)}),
		server.WithLimit(&limit.Option{MaxConnections: conf.ServerMaxConnections, MaxQPS: conf.ServerMaxQPS}), // limit
		server.WithMuxTransport(),                       // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite())) // tracer)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
