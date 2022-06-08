package main

import (
	"context"
	"ihome/service/captcha/conf"
	"ihome/service/user/kitex_gen"
	"ihome/service/user/model"
	"ihome/service/utils"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// SendSMS implements the UserServiceImpl interface.
func (s *UserServiceImpl) SendSMS(ctx context.Context, req *kitex_gen.SMSRequest) (resp *kitex_gen.Response, err error) {
	utils.NewLog().Info("SendSMS...", req.Phone, req.ImgCode, req.Uuid)
	response := model.SendSMSCode(req.Phone, req.ImgCode, req.Uuid)
	utils.NewLog().Info("SendSMS...", response)
	if utils.RECODE_OK != response.Errno {
		utils.NewLog().Error("SendSMSCode error:", response)
		return &response, nil
	}
	//保存短信验证码到redis
	response = model.SaveSMSCode(req.Phone, conf.PhoneCode)
	utils.NewLog().Info("SaveSMSCode...", response)
	if utils.RECODE_OK != response.Errno {
		utils.NewLog().Error("SaveSMSCode error:", response)
		return &response, nil
	}
	return &response, nil

}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *kitex_gen.RegRequest) (resp *kitex_gen.Response, err error) {

	utils.NewLog().Info("Register...", req.Phone+":"+req.Password+":"+req.SmsCode)
	response := model.CheckSMSCode(req.Phone, req.SmsCode)
	if utils.RECODE_OK != response.Errno {
		utils.NewLog().Error("CheckSMSCode error", response)
		return &response, nil
	}
	response = model.Register(req.Phone, req.Password)

	if utils.RECODE_OK != response.Errno {
		utils.NewLog().Error("Register error", response)
		return &response, nil
	}
	response = utils.UserResponse(utils.RECODE_OK)
	return &response, nil
}
