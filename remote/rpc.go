package remote

import (
	"context"
	"errors"
	"ihome/conf"
	captcha_kitex_gen "ihome/service/captcha/kitex_gen"
	"ihome/service/captcha/kitex_gen/captchaservice"
	house_kitex_gen "ihome/service/house/kitex_gen"
	"ihome/service/house/kitex_gen/houseservice"
	order_kitex_gen "ihome/service/order/kitex_gen"
	"ihome/service/order/kitex_gen/orderservice"
	user_kitex_gen "ihome/service/user/kitex_gen"
	"ihome/service/user/kitex_gen/userservice"
	"ihome/service/utils"
)

var Ctx, _ = context.WithTimeout(context.Background(), conf.RPCTimeOut)

func RPC(ctx context.Context, serviceName string, req interface{}) (interface{}, error) {
	result := GetService(&ctx, serviceName)
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
	case conf.OrderServerIndex:
		response, err = handlerOrderService(ctx, result, req)
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

func handlerHouseService(ctx context.Context, result interface{}, req interface{}) (interface{}, error) {
	service := result.(houseservice.Client)
	var response interface{}
	var err error
	switch req.(type) {
	case house_kitex_gen.AreaRequest:
		request := req.(house_kitex_gen.AreaRequest)
		response, err = service.GetArea(ctx, &request)
	case house_kitex_gen.PubHouseRequest:
		request := req.(house_kitex_gen.PubHouseRequest)
		response, err = service.PubHouse(ctx, &request)
	case house_kitex_gen.GetUserHouseRequest:
		request := req.(house_kitex_gen.GetUserHouseRequest)
		response, err = service.GetUserHouse(ctx, &request)
	case house_kitex_gen.UploadHouseImgReq:
		request := req.(house_kitex_gen.UploadHouseImgReq)
		response, err = service.UploadHouseImg(ctx, &request)
	case house_kitex_gen.GetHouseDetailReg:
		request := req.(house_kitex_gen.GetHouseDetailReg)
		response, err = service.GetHouseDetail(ctx, &request)
	case house_kitex_gen.HouseSearchReq:
		request := req.(house_kitex_gen.HouseSearchReq)
		response, err = service.SearchHouse(ctx, &request)
	case house_kitex_gen.HouseHomeIndexReg:
		request := req.(house_kitex_gen.HouseHomeIndexReg)
		response, err = service.HouseHomeIndex(ctx, &request)
	case house_kitex_gen.GetHouseInfoReq:
		request := req.(house_kitex_gen.GetHouseInfoReq)
		response, err = service.GetHouseInfo(ctx, &request)
	default:
		err = errors.New("handlerHouseService error")
	}
	return response, err
}

func handlerOrderService(ctx context.Context, result interface{}, req interface{}) (interface{}, error) {
	service := result.(orderservice.Client)
	var response interface{}
	var err error
	switch req.(type) {
	case order_kitex_gen.PostOrderReg:
		request := req.(order_kitex_gen.PostOrderReg)
		response, err = service.PostOrder(ctx, &request)

	default:
		err = errors.New("handlerHouseService error")
	}
	return response, err
}

func handlerUserService(ctx context.Context, result interface{}, req interface{}) (interface{}, error) {
	service := result.(userservice.Client)
	utils.NewLog().Info("userService.Client", service)
	var response interface{}
	var err error
	//utils.NewLog().Info("req type", req)
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
		utils.NewLog().Debug("SessionAuthRequest")
		request := req.(user_kitex_gen.SessionAuthRequest)
		response, err = service.SessionAuth(ctx, &request)
	case user_kitex_gen.SessionRequest:
		utils.NewLog().Debug("SessionRequest")
		request := req.(user_kitex_gen.SessionRequest)
		response, err = service.GetSessionInfo(ctx, &request)
	case user_kitex_gen.SessionDeleteRequest:
		utils.NewLog().Debug("SessionDeleteRequest")
		request := req.(user_kitex_gen.SessionDeleteRequest)
		response, err = service.DeleteSession(ctx, &request)
	case user_kitex_gen.GetUserRequest:
		utils.NewLog().Debug("GetUserRequest")
		request := req.(user_kitex_gen.GetUserRequest)
		response, err = service.GetUserInfo(ctx, &request)
	case user_kitex_gen.UpdateUserRequest:
		utils.NewLog().Debug("UpdateUserRequest")
		request := req.(user_kitex_gen.UpdateUserRequest)
		response, err = service.UpdateUserInfo(ctx, &request)
	case user_kitex_gen.UploadImgRequest:
		utils.NewLog().Debug("UpdateUserRequest")
		request := req.(user_kitex_gen.UploadImgRequest)
		response, err = service.UploadImg(ctx, &request)
	default:
		utils.NewLog().Info("handlerUserService default")
		err = errors.New("handlerUserService error")
	}
	return response, err

}

func handlerCaptchaService(ctx context.Context, result interface{}, req interface{}) (interface{}, error) {
	//验证码服务
	var response interface{}
	var err error
	service := result.(captchaservice.Client)
	request := req.(captcha_kitex_gen.Request)
	response, err = service.GetCaptcha(ctx, &request)
	if err != nil {
		utils.NewLog().Error("service.GetCaptcha success")

		return response, err
	}
	return response, nil

}
