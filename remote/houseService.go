package remote

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"ihome/conf"
	"ihome/service/house/kitex_gen/houseservice"
	"ihome/service/utils"
	"strconv"
	"time"
)

type HouseServiceClient struct {
	Service houseservice.Client
	Err     error
}

var HouseService *HouseServiceClient

func init() {
	HouseService = &HouseServiceClient{}
	service, err := houseservice.NewClient(conf.HouseServiceIndex,
		client.WithHostPorts(conf.HouseServerIp+":"+strconv.Itoa(conf.HouseServerPort)),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()))   // tracer)
	if err != nil {
		utils.NewLog().Error("HouseService init error")
	}
	HouseService.Err = err
	HouseService.Service = service
}
