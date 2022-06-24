package controller

import (
	"encoding/json"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"ihome/remote"
	captcha_kitex_gen "ihome/service/captcha/kitex_gen"
	"ihome/web/conf"
	"ihome/web/utils"
	"image/png"
	"strconv"
)

//GetImageCd 获取验证码
func GetImageCd(ctx *gin.Context) {

	uuid := ctx.Param("uuid")
	utils.NewLog().Info("", conf.CaptchaServerIp+":"+strconv.Itoa(conf.CaptchaServerPort))
	req := captcha_kitex_gen.Request{Uuid: uuid}
	resp, err := remote.RPC(ctx, conf.CaptchaServiceIndex, req)
	if err != nil {
		utils.NewLog().Info("GetCaptcha error ...", err)
		return
	}
	var img captcha.Image
	//json反序列化
	json.Unmarshal(resp.(*captcha_kitex_gen.Response).Img, &img)
	png.Encode(ctx.Writer, img)
	utils.NewLog().Info("uuid:", uuid)
}
