package controller

import (
	"github.com/gin-gonic/gin"
	captcha_kitex_gen "ihome/service/captcha/kitex_gen"
	"ihome/web/conf"
	"ihome/web/remote"
	"ihome/web/utils"
	"strconv"
)

//GetImageCd 获取验证码
func GetImageCd(ctx *gin.Context) {

	uuid := ctx.Param("uuid")
	utils.NewLog().Info("", conf.CaptchaServerIp+":"+strconv.Itoa(conf.CaptchaServerPort))
	req := captcha_kitex_gen.Request{Uuid: uuid}
	_, err := remote.RPC(ctx, conf.CaptchaServiceIndex, req)
	if err != nil {
		utils.NewLog().Info("GetCaptcha error ...", err)
		return
	}
	utils.NewLog().Info("uuid:", uuid)
}
