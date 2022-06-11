package utils

import (
	"errors"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"ihome/service/captcha/kitex_gen/captchaservice"
	"ihome/service/house/kitex_gen/houseservice"
	"ihome/service/user/kitex_gen/userservice"
	"ihome/service/utils"
	"ihome/web/conf"
	"strconv"
)

func GetService(ctx *gin.Context, serviceName string) interface{} {
	var service interface{}
	var err error
	utils.NewLog().Info("serviceName:", serviceName)
	switch serviceName {
	case conf.CaptchaServiceIndex:
		service, err = captchaservice.NewClient(conf.CaptchaServiceIndex,
			client.WithHostPorts(conf.CaptchaServerIp+":"+strconv.Itoa(conf.CaptchaServerPort)))
	case conf.UserServiceIndex:
		service, err = userservice.NewClient(conf.UserServiceIndex,
			client.WithHostPorts(conf.UserServerIp+":"+strconv.Itoa(conf.UserServerPort)))
	case conf.HouseServiceIndex:
		utils.NewLog().Info("start GetService HouseServiceIndex...")
		service, err = houseservice.NewClient(conf.HouseServiceIndex,
			client.WithHostPorts(conf.HouseServerIp+":"+strconv.Itoa(conf.HouseServerPort)))
	default:
		utils.NewLog().Info("default")
		service = nil
		err = errors.New("no service error")
	}
	if err != nil {
		NewLog().Error("get service error:", err)
		//ctx.JSON(http.StatusOK, Response(RECODE_SERVERERR, nil))
		return nil

	}
	return service

}
