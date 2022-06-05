package main

import (
	"context"
	"ihome/service/captcha/conf"
	"ihome/service/user/kitex_gen"
	"ihome/service/user/model"
	"ihome/service/user/utils"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// SendSMS implements the UserServiceImpl interface.
func (s *UserServiceImpl) SendSMS(ctx context.Context, req *kitex_gen.SMSRequest) (resp *kitex_gen.Response, err error) {
	utils.NewLog().Info("SendSMS...", req.Phone, req.ImgCode, req.Uuid)
	errs := model.SendSMSCode(req.Phone, req.ImgCode, req.Uuid)
	utils.NewLog().Info("SendSMS...", errs)
	response := &kitex_gen.Response{}
	if errs != nil {
		response.Errno = utils.RECODE_SMSERR
		response.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		return response, errs
	}
	//保存短信验证码到redis
	errs = model.SaveSMSCode(req.Phone, conf.PhoneCode)
	utils.NewLog().Info("SaveSMSCode...", errs)
	if errs != nil {
		response.Errno = utils.RECODE_SMSERR
		response.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		return response, errs
	}
	response.Errno = utils.RECODE_OK
	response.Errmsg = utils.RecodeText(utils.RECODE_OK)
	return response, nil

}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *kitex_gen.SMSRequest) (resp *kitex_gen.Response, err error) {
	// TODO: Your code here...

	return
}
