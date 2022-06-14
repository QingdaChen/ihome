package main

import (
	"context"
	"ihome/service/user/conf"
	"ihome/service/user/handler"
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
	response = utils.UserResponse(utils.RECODE_OK, nil)
	return &response, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *kitex_gen.LoginRequest) (resp *kitex_gen.Response, err error) {
	utils.NewLog().Debug("login :", req.Phone, req.Password)
	//检查数据库
	loginResp := model.Login(req.Phone, req.Password)
	//utils.NewLog().Info("login info:", loginResp)
	if utils.RECODE_OK != loginResp.Errno {
		utils.NewLog().Info("mysql login error:", loginResp)
		return &loginResp, nil
	}
	//保存session
	sessionResp := model.SaveRedisSession(loginResp.Data)
	if utils.RECODE_OK != sessionResp.Errno {
		utils.NewLog().Info("SaveRedisSession error:", sessionResp)
	}
	utils.NewLog().Info("sessionResp:", string(sessionResp.Data))
	return &sessionResp, nil

}

// SessionAuth implements the UserServiceImpl interface.
func (s *UserServiceImpl) SessionAuth(ctx context.Context, req *kitex_gen.SessionAuthRequest) (resp *kitex_gen.Response, err error) {
	utils.NewLog().Debug("SessionAuth :", req.SessionId)
	checkResp := model.CheckRedisSession(req.SessionId)
	utils.NewLog().Info("checkResp:", checkResp)
	return &checkResp, nil
}

// GetSessionInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetSessionInfo(ctx context.Context, req *kitex_gen.SessionRequest) (resp *kitex_gen.Response, err error) {
	utils.NewLog().Debug("GetSessionInfo start")
	//调用redis查询session信息
	sessionResp := model.GetSessionInfo(req.SessionId)
	//utils.NewLog().Info("GetSessionInfo:", sessionResp)

	return &sessionResp, nil

}

// DeleteSession implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteSession(ctx context.Context, req *kitex_gen.SessionDeleteRequest) (resp *kitex_gen.Response, err error) {
	utils.NewLog().Debug("DeleteSession start")
	//调用redis删除session
	deleteResp := model.DeleteKey(conf.SessionLoginIndex + "_" + req.SessionId)
	utils.NewLog().Info("GetSessionInfo:", deleteResp)
	return &deleteResp, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *kitex_gen.GetUserRequest) (resp *kitex_gen.Response, err error) {
	utils.NewLog().Debug("GetUserInfo start")
	//查询用户信息
	userResp := handler.GetUserInfoByHandler(req.SessionId)
	if utils.RECODE_OK != userResp.Errno {
		utils.NewLog().Error("getUserInfoByHandler error", userResp)
	}
	return &userResp, nil
}

// UpdateUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserInfo(ctx context.Context, req *kitex_gen.UpdateUserRequest) (resp *kitex_gen.Response, err error) {
	utils.NewLog().Debug("UpdateUserInfo start")
	//使用分布式锁
	//model.InitLock(conf.RedisLockKey + "-user-" + req.SessionId)
	//model.Lock.Lock()
	if err != nil {
		utils.NewLog().Error("get redisLock fail:", err)
		response := utils.UserResponse(utils.RECODE_SERVERERR, nil)
		return &response, nil
	}
	//defer model.TryRelease()
	//更新用户信息
	userInfoResp := handler.UpdateUserInfoByHandler(req.SessionId, "name", req.UpdateName)
	if utils.RECODE_OK != userInfoResp.Errno {
		utils.NewLog().Info("UpdateUserInfoByHandler error:", userInfoResp)
		return &userInfoResp, nil
	}
	//在更新session
	sessionResp := model.UpdateSessionInfo(req.SessionId, req.UpdateName)
	if utils.RECODE_OK != sessionResp.Errno {
		//TODO 发送消息错误兜底
		utils.NewLog().Info("UpdateSessionInfo failed")
	}
	return &sessionResp, nil
}

// UploadImg implements the UserServiceImpl interface.
func (s *UserServiceImpl) UploadImg(ctx context.Context, req *kitex_gen.UploadImgRequest) (resp *kitex_gen.Response, err error) {
	utils.NewLog().Debug("UploadImg start....")
	//使用分布式锁，TODO 改进
	//model.InitLock(conf.RedisLockKey + "-user-" + req.SessionId)
	//model.Lock.Lock()
	//defer model.TryRelease()
	//先上传
	//utils.NewLog().Info("base64:  ", req.ImgBase64)

	uploadResp := model.FastDfsClient.UploadImg(req.ImgBase64, req.FileType)
	if utils.RECODE_OK != uploadResp.Errno {
		//TODO 存数据库记录
		utils.NewLog().Info("UploadImg error", uploadResp)
		return &uploadResp, nil
	}
	utils.NewLog().Info("uploadResp:", string(uploadResp.Data))
	//更新缓存
	//model.UpdateSessionInfo()
	//查头像原url
	userResp := handler.GetUserInfoByHandler(req.SessionId)
	if utils.RECODE_OK != userResp.Errno {
		//TODO 存数据库记录
		utils.NewLog().Info("GetUserInfoByHandler error", userResp)
		return &userResp, nil
	}
	//成功就更新数据库和缓存
	updateResp := handler.UpdateUserInfoByHandler(req.SessionId, "avatar_url", (string)(uploadResp.Data))
	if utils.RECODE_OK != updateResp.Errno {
		//TODO 存数据库记录
		utils.NewLog().Info("UpdateUserInfoByHandler error", updateResp)
		return &updateResp, nil
	}
	//都成功启动协程删除原头像
	handler.DeleteFileByHandler(userResp.Data)
	//拼接头像地址
	//fileName := fmt.Sprintf("%s/%s", conf.NginxUrl, (string)(uploadResp.Data))
	//utils.NewLog().Debug("fileName :", fileName)
	//uploadResp.Data = []byte(fileName)
	return &uploadResp, nil
}
