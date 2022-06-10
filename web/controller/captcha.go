package controller

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	captcha_kitex_gen "ihome/service/captcha/kitex_gen"
	"ihome/service/captcha/kitex_gen/captchaservice"
	"ihome/web/conf"
	"ihome/web/utils"
	"image/png"
	"net/http"
	"strconv"
)

//GetImageCd 获取验证码
func GetImageCd(ctx *gin.Context) {

	uuid := ctx.Param("uuid")
	result, resp := utils.GetService(conf.CaptchaServiceIndex)
	if utils.RECODE_OK != resp[conf.ErrorNoIndex].(string) {
		ctx.JSON(http.StatusOK, resp)
		return
	}
	service := result.(captchaservice.Client)
	utils.NewLog().Info("", conf.CaptchaServerIp+":"+strconv.Itoa(conf.CaptchaServerPort))
	utils.NewLog().Info("client ...", service)

	req := &captcha_kitex_gen.Request{Uuid: uuid}
	response, err2 := service.GetCaptcha(context.Background(), req)
	if err2 != nil {
		utils.NewLog().Info("GetCaptcha error ...", err2)
	}
	var img captcha.Image
	err2 = json.Unmarshal(response.Img, &img)
	if err2 != nil {
		utils.NewLog().Error("json.Unmarshal error:", err2)
		utils.Resp(resp, utils.RECODE_SERVERERR)
		ctx.JSON(http.StatusOK, resp)
	}
	png.Encode(ctx.Writer, img)
	utils.NewLog().Info("uuid:", uuid)

}
