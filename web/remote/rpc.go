package remote

import (
	"encoding/json"
	"errors"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	captcha_kitex_gen "ihome/service/captcha/kitex_gen"
	"ihome/service/captcha/kitex_gen/captchaservice"
	house_kitex_gen "ihome/service/house/kitex_gen"
	"ihome/service/house/kitex_gen/houseservice"
	user_kitex_gen "ihome/service/user/kitex_gen"
	"ihome/service/user/kitex_gen/userservice"
	"ihome/web/conf"
	"ihome/web/utils"
	"image/png"
)

func RPC(ctx *gin.Context, serviceName string, req interface{}) (interface{}, error) {
	result := utils.GetService(ctx, serviceName)
	utils.NewLog().Info("GetService:", result)
	if result == nil {
		return nil, errors.New("utils.GetService error")
	}
	var response interface{}
	var err error
	switch serviceName {
	case conf.CaptchaServiceIndex:
		response, err = handlerCaptchaService(ctx, result, req)
	case conf.UserServiceIndex:
		utils.NewLog().Info("UserServiceIndex..")
		response, err = handlerUserService(ctx, result, req)
	case conf.HouseServiceIndex:
		response, err = handlerHouseService(ctx, result, req)
	default:
		//ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return utils.Response(utils.RECODE_SERVERERR, nil), errors.New("service rpc error")
	}
	if err != nil {
		utils.NewLog().Error("rpc error:", err)
		//ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return utils.Response(utils.RECODE_SERVERERR, nil), errors.New("service rpc error")
	}
	//ctx.JSON(http.StatusOK, response)
	return response, nil
}

func handlerHouseService(ctx *gin.Context, result interface{}, req interface{}) (interface{}, error) {
	service := result.(houseservice.Client)
	var response interface{}
	var err error
	switch req.(type) {
	case house_kitex_gen.AreaRequest:
		request := req.(house_kitex_gen.AreaRequest)
		response, err = service.GetArea(ctx, &request)
	default:
		err = errors.New("handlerHouseService error")
	}
	return response, err
}

func handlerUserService(ctx *gin.Context, result interface{}, req interface{}) (interface{}, error) {
	service := result.(userservice.Client)
	utils.NewLog().Info("userService.Client", service)
	var response interface{}
	var err error
	utils.NewLog().Infof("req type:%v", req)
	switch req.(type) {
	case user_kitex_gen.SMSRequest:
		request := req.(user_kitex_gen.SMSRequest)
		response, err = service.SendSMS(ctx, &request)
	case user_kitex_gen.RegRequest:
		request := req.(user_kitex_gen.RegRequest)
		response, err = service.Register(ctx, &request)
	case user_kitex_gen.LoginRequest:
		request := req.(user_kitex_gen.LoginRequest)
		response, err = service.Login(ctx, &request)
	case user_kitex_gen.SessionAuthRequest:
		utils.NewLog().Info("user_kitex_gen.SessionAuthRequest")
		request := req.(user_kitex_gen.SessionAuthRequest)
		response, err = service.SessionAuth(ctx, &request)
		utils.NewLog().Info("user_kitex_gen.SessionAuthRequest:", response)
	default:
		utils.NewLog().Info("handlerUserService default")
		err = errors.New("handlerUserService error")
	}
	return response, err

}

func handlerCaptchaService(ctx *gin.Context, result interface{}, req interface{}) (interface{}, error) {
	//验证码服务
	var response interface{}
	var err error
	service := result.(captchaservice.Client)
	request := req.(captcha_kitex_gen.Request)
	response, err = service.GetCaptcha(ctx, &request)
	var img captcha.Image
	//json反序列化
	err2 := json.Unmarshal(response.(captcha_kitex_gen.Response).Img, &img)
	if err2 == nil {
		utils.NewLog().Error("json.Unmarshal success")
		//写入验证码图像
		png.Encode(ctx.Writer, img)
		return response, nil
	}
	err = errors.New("service.GetCaptcha")
	return response, err

}
