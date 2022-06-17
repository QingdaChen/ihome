package remote

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"ihome/service/house/conf"
	"ihome/service/user/kitex_gen/userservice"
	"ihome/service/utils"
	"strconv"
	"time"
)

type UserServiceClient struct {
	Service userservice.Client
	Err     error
}

var UserService *UserServiceClient

func init() {
	UserService = &UserServiceClient{}
	service, err := userservice.NewClient(conf.UserServiceIndex,
		client.WithHostPorts(conf.UserServerIp+":"+strconv.Itoa(conf.UserServerPort)),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(10*time.Second),             // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()))   // tracer)
	if err != nil {
		utils.NewLog().Error("UserServiceClient init error")
	}
	UserService.Err = err
	UserService.Service = service
}
