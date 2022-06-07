package controller

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	captcha_kitex_gen "ihome/service/captcha/kitex_gen"
	"ihome/service/captcha/kitex_gen/captchaservice"
	user_kitex_gen "ihome/service/user/kitex_gen"
	"ihome/service/user/kitex_gen/userservice"
	"ihome/web/conf"
	"ihome/web/model"

	"ihome/web/utils"
	"image/png"
	"net/http"
	"strconv"
)

func GetSession(ctx *gin.Context) {

	resp := map[string]string{}
	utils.Resp(resp, utils.RECODE_SESSIONERR)
	ctx.JSON(http.StatusOK, resp)

}

func GetImageCd(ctx *gin.Context) {

	initCaptcha(ctx)

}

// GetSMSCd http://xx.com/api/v1.0/smscode/111?text=248484&id=9cd8faa9-5653-4f7c-b653-0a58a8a98c81
func GetSMSCd(ctx *gin.Context) {
	//TODO 接口防刷
	phone := ctx.Param("phone")
	imgCode := ctx.Query("text")
	uuid := ctx.Query("id")

	client, err := userservice.NewClient("userService",
		client.WithHostPorts(conf.UserServerIp+":"+strconv.Itoa(conf.UserServerPort)),
	)
	utils.NewLog().Info("GetSMSCd..." + phone + ":" + imgCode + ":" + uuid)
	resp := map[string]string{}
	//连接错误
	if err != nil {
		utils.Resp(resp, utils.RECODE_SERVERERR)
		ctx.JSON(http.StatusOK, resp)
		return
	}
	//发送短信
	req := &user_kitex_gen.SMSRequest{Phone: phone, ImgCode: imgCode, Uuid: uuid}
	response, _ := client.SendSMS(context.Background(), req)
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
	var service userservice.Client
	service, err = userservice.NewClient("userService",
		client.WithHostPorts(conf.UserServerIp+":"+strconv.Itoa(conf.UserServerPort)),
	)
	resp := map[string]string{}
	//连接错误
	if err != nil {
		utils.Resp(resp, utils.RECODE_SERVERERR)
		ctx.JSON(http.StatusOK, resp)
		return
	}
	req := &user_kitex_gen.RegRequest{Phone: register.Phone,
		Password: register.Password, SmsCode: register.SmsCode}
	response, _ := service.Register(ctx, req)
	utils.NewLog().Info("service.Register:", response)
	utils.Resp(resp, response.Errno)
	ctx.JSON(http.StatusOK, resp)
	return

}

//获取验证码
func initCaptcha(ctx *gin.Context) {

	uuid := ctx.Param("uuid")
	client, err := captchaservice.NewClient("captchaService",
		client.WithHostPorts(conf.CaptchaServerIp+":"+strconv.Itoa(conf.CaptchaServerPort)),
		//client.WithResolver(dns.NewDNSResolver()),
	)
	utils.NewLog().Info("", conf.CaptchaServerIp+":"+strconv.Itoa(conf.CaptchaServerPort))
	utils.NewLog().Info("client ...", client)
	if err != nil {
		utils.NewLog().Info("captchaService error...", err)
	}
	req := &captcha_kitex_gen.Request{Uuid: uuid}
	resp, err2 := client.GetCaptcha(context.Background(), req)
	if err2 != nil {
		utils.NewLog().Info("GetCaptcha error ...", err2)
	}
	var img captcha.Image
	json.Unmarshal(resp.Img, &img)
	png.Encode(ctx.Writer, img)
	utils.NewLog().Info("uuid:", uuid)
}
