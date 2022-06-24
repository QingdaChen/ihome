package remote

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"ihome/conf"
	"ihome/service/order/kitex_gen/orderservice"
	"ihome/service/utils"
	"strconv"
	"time"
)

type OrderServiceClient struct {
	Service orderservice.Client
	Err     error
}

var OrderService *OrderServiceClient

func init() {
	OrderService = &OrderServiceClient{}
	service, err := orderservice.NewClient(conf.OrderServerIndex,
		client.WithHostPorts(conf.OrderServerIp+":"+strconv.Itoa(conf.OrderServerPort)),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()))   // tracer)
	if err != nil {
		utils.NewLog().Error("HouseService init error")
	}
	OrderService.Err = err
	OrderService.Service = service
}
