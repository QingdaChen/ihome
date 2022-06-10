package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	user_kitex_gen "ihome/service/user/kitex_gen"
	"ihome/service/user/kitex_gen/userservice"
	"ihome/web/conf"
	"ihome/web/model"
	"ihome/web/utils"
	"net/http"
)

func GetSession(ctx *gin.Context) {

	resp := make(map[string]interface{})
	utils.Resp(resp, utils.RECODE_SESSIONERR)
	ctx.JSON(http.StatusOK, resp)

}

// GetSMSCd http://xx.com/api/v1.0/smscode/111?text=248484&id=9cd8faa9-5653-4f7c-b653-0a58a8a98c81
func GetSMSCd(ctx *gin.Context) {

	phone := ctx.Param("phone")
	imgCode := ctx.Query("text")
	uuid := ctx.Query("id")
	//接口防刷
	ip, _ := ctx.RemoteIP()
	utils.NewLog().Info("remote ip phone:", ip, phone)
	resp := make(map[string]interface{})
	res1, _ := utils.SMSCache.Get(string(ip))
	res2, _ := utils.SMSCache.Get(phone)
	resIp := string(res1)
	resPhone := string(res2)
	utils.NewLog().Info("InterfaceCache:", resIp, resPhone)
	if (resIp != "") || (resPhone != "") {
		utils.NewLog().Info("cache:", resIp, resPhone)
		utils.Resp(resp, utils.RECODE_REQFRE)
		ctx.JSON(http.StatusOK, resp)
		return
	}
	utils.SMSCache.Set(string(ip), []byte("1"))
	utils.SMSCache.Set(phone, []byte("1"))
	//utils.NewLog().Info("cache:", utils.InterfaceCache)

	//连接服务
	result, resp := utils.GetService(conf.UserServiceIndex)
	if utils.RECODE_OK != resp[conf.ErrorNoIndex].(string) {
		ctx.JSON(http.StatusOK, resp)
		return
	}
	service := result.(userservice.Client)

	utils.NewLog().Info("GetSMSCd..." + phone + ":" + imgCode + ":" + uuid)

	//发送短信
	req := &user_kitex_gen.SMSRequest{Phone: phone, ImgCode: imgCode, Uuid: uuid}
	response, err := service.SendSMS(context.Background(), req)
	utils.NewLog().Info("SendSMS...", err)
	utils.Resp(resp, response.Errno)
	ctx.JSON(http.StatusOK, resp)
	return

}

// Register 注册
func Register(ctx *gin.Context) {

	var register model.RegisterRequest
	err := ctx.Bind(&register)
	utils.NewLog().Info("register:", register)
	if err != nil {
		utils.NewLog().Error("Register bind error:", err)
		return
	}
	//连接服务
	result, resp := utils.GetService(conf.UserServiceIndex)
	if utils.RECODE_OK != resp[conf.ErrorNoIndex].(string) {
		ctx.JSON(http.StatusOK, resp)
		return
	}
	service := result.(userservice.Client)
	req := &user_kitex_gen.RegRequest{Phone: register.Phone,
		Password: register.Password, SmsCode: register.SmsCode}
	response, _ := service.Register(ctx, req)
	utils.NewLog().Info("service.Register:", response)
	utils.Resp(resp, response.Errno)
	ctx.JSON(http.StatusOK, resp)
	return

}

//PostLogin 登录
func PostLogin(ctx *gin.Context) {
	var loginUser model.LoginRequest
	err := ctx.Bind(&loginUser)
	if err != nil {
		utils.NewLog().Error("Bind error:", err)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return
	}
	//连接服务
	result, resp := utils.GetService(conf.UserServiceIndex)
	if utils.RECODE_OK != resp[conf.ErrorNoIndex].(string) {
		ctx.JSON(http.StatusOK, resp)
		return
	}
	service := result.(userservice.Client)
	req := &user_kitex_gen.LoginRequest{Phone: loginUser.Phone, Password: loginUser.Password}
	response, err2 := service.Login(ctx, req)
	if err2 != nil {
		utils.NewLog().Error("LoginRequest error:", err2)
		resp = utils.Response(utils.RECODE_SERVERERR, nil)
		ctx.JSON(http.StatusOK, resp)
		return
	}
	if utils.RECODE_OK == response.Errno {
		//如果登录成功就保存cookie
		utils.NewLog().Info("login response:", string(response.Data))
		ctx.SetCookie(conf.LoginCookieName, string(response.Data),
			conf.LoginCookieTimeOut, "/", conf.UserServerIp, false, true)
	}
	resp = utils.Response(response.Errno, nil)
	ctx.JSON(http.StatusOK, resp)
	return
}
