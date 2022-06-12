package remote

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"ihome/service/captcha/kitex_gen/captchaservice"
	"ihome/web/conf"
	"ihome/web/utils"
	"strconv"
	"time"
)

type CaptchaServiceClient struct {
	Service captchaservice.Client
	Err     error
}

var CaptchaService *CaptchaServiceClient

func init() {
	CaptchaService = &CaptchaServiceClient{}
	service, err := captchaservice.NewClient(conf.CaptchaServiceIndex,
		client.WithHostPorts(conf.CaptchaServerIp+":"+strconv.Itoa(conf.CaptchaServerPort)),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()))   // tracer

	if err != nil {
		utils.NewLog().Error("CaptchaService init error")
	}
	CaptchaService.Err = err
	CaptchaService.Service = service
}
