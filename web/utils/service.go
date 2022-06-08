package utils

import (
	"errors"
	"github.com/cloudwego/kitex/client"
	"ihome/service/captcha/kitex_gen/captchaservice"
	"ihome/service/house/kitex_gen/houseservice"
	"ihome/service/user/kitex_gen/userservice"
	"ihome/service/utils"
	"ihome/web/conf"
	"strconv"
)

func GetService(serviceName string) (interface{}, map[string]interface{}) {
	resp := make(map[string]interface{})
	var service interface{}
	var err error
	switch serviceName {
	case conf.CaptchaServiceIndex:
		service, err = captchaservice.NewClient(conf.CaptchaServiceIndex,
			client.WithHostPorts(conf.CaptchaServerIp+":"+strconv.Itoa(conf.CaptchaServerPort)))
	case conf.UserServiceIndex:
		service, err = userservice.NewClient(conf.UserServiceIndex,
			client.WithHostPorts(conf.UserServerIp+":"+strconv.Itoa(conf.UserServerPort)))
	case conf.HouseServiceIndex:
		service, err = houseservice.NewClient(conf.HouseServiceIndex,
			client.WithHostPorts(conf.HouseServerIp+":"+strconv.Itoa(conf.HouseServerPort)))
	default:
		service = nil
		err = errors.New("no service error")
	}
	if err != nil {
		NewLog().Error("get service error:", err)
		Resp(resp, utils.RECODE_SERVERERR)
		return nil, resp
	}
	Resp(resp, utils.RECODE_OK)
	return service, resp

}
