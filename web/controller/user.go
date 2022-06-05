package controller

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	captcha_kitex_gen "ihome/service/captcha/kitex_gen"
	"ihome/service/captcha/kitex_gen/captchaservice"
	user_kitex_gen "ihome/service/user/kitex_gen"
	"ihome/service/user/kitex_gen/userservice"
	"ihome/web/conf"

	"ihome/web/utils"
	"image/png"
	"net/http"
	"strconv"
)

func GetSession(ctx *gin.Context) {

	resp := map[string]string{}
	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	ctx.JSON(http.StatusOK, resp)

}

func GetImageCd(ctx *gin.Context) {

	initCaptcha(ctx)

}

// GetSMSCd http://xx.com/api/v1.0/smscode/111?text=248484&id=9cd8faa9-5653-4f7c-b653-0a58a8a98c81
func GetSMSCd(ctx *gin.Context) {
	phone := ctx.Param("phone")
	imgCode := ctx.Query("text")
	uuid := ctx.Query("id")

	client, err := userservice.NewClient("userService",
		client.WithHostPorts(conf.UserServerIp+":"+strconv.Itoa(conf.UserServerPort)),
	)
	utils.NewLog().Info("GetSMSCd..." + phone + ":" + imgCode + ":" + uuid)
	resp := map[string]string{}
	//连接...
	if err != nil {

		resp["errno"] = utils.RECODE_SMSERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SMSERR)
		ctx.JSON(http.StatusOK, resp)
		return

	}
	req := &user_kitex_gen.SMSRequest{Phone: phone, ImgCode: imgCode, Uuid: uuid}

	_, err = client.SendSMS(context.Background(), req)
	utils.NewLog().Info("SendSMS...", err)
	if err != nil {
		resp["errno"] = utils.RECODE_SMSERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SMSERR)
		ctx.JSON(http.StatusOK, resp)
		return
	}

	//比对和保存成功...
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	ctx.JSON(http.StatusOK, resp)

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
		logrus.Info("captchaService error...", err)
	}
	req := &captcha_kitex_gen.Request{Uuid: uuid}
	resp, err2 := client.GetCaptcha(context.Background(), req)
	if err2 != nil {
		logrus.Info("GetCaptcha error ...", err2)
	}
	var img captcha.Image
	json.Unmarshal(resp.Img, &img)
	png.Encode(ctx.Writer, img)
	utils.NewLog().Info("uuid:", uuid)
}
