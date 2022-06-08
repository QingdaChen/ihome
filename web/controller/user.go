package controller

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	captcha_kitex_gen "ihome/service/captcha/kitex_gen"
	"ihome/service/captcha/kitex_gen/captchaservice"
	house_kitex_gen "ihome/service/house/kitex_gen"
	"ihome/service/house/kitex_gen/houseservice"
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

	resp := make(map[string]interface{})
	utils.Resp(resp, utils.RECODE_SESSIONERR)
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

	client, err := userservice.NewClient("userService",
		client.WithHostPorts(conf.UserServerIp+":"+strconv.Itoa(conf.UserServerPort)),
	)
	utils.NewLog().Info("GetSMSCd..." + phone + ":" + imgCode + ":" + uuid)

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
	resp := make(map[string]interface{})
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

//GetAreas 获取地区信息
func GetAreas(ctx *gin.Context) {
	//先走本地缓存
	resp := make(map[string]interface{})
	resp["data"] = ""
	//缓存中查
	cacheAreas, _ := utils.AreasCache.Get(conf.HouseAreasCacheIndex)
	var areas []model.Area
	if cacheAreas != nil {
		utils.NewLog().Info("cacheAreas:", string(cacheAreas))
		utils.Resp(resp, utils.RECODE_OK)
		err := json.Unmarshal(cacheAreas, &areas)
		if err != nil {
			utils.NewLog().Error("json.Unmarshal error:", err)
			utils.Resp(resp, utils.RECODE_SERVERERR)
			ctx.JSON(http.StatusOK, resp)
			return
		}
		utils.NewLog().Info("cache areas:", areas)
		//直接返回
		resp["data"] = areas
		ctx.JSON(http.StatusOK, resp)
		return
	}
	//查不到远程请求house服务查询
	service, err := houseservice.NewClient("houseservice",
		client.WithHostPorts(conf.HouseServerIp+":"+strconv.Itoa(conf.HouseServerPort)),
	)
	//连接错误
	if err != nil {
		utils.NewLog().Info("houseservice connect error", err)
		utils.Resp(resp, utils.RECODE_SERVERERR)
		ctx.JSON(http.StatusOK, resp)
		return
	}
	req := &house_kitex_gen.AreaRequest{}
	response, _ := service.GetArea(ctx, req)
	utils.Resp(resp, response.Errno)
	if utils.RECODE_OK != response.Errno {
		resp["data"] = ""
		ctx.JSON(http.StatusOK, resp)
		return
	}
	//反序列化
	err = json.Unmarshal(response.Data, &areas)
	if err != nil {
		utils.NewLog().Info("json.Unmarshal(response.Data, &areas) error", err)
		utils.Resp(resp, utils.RECODE_SERVERERR)
		ctx.JSON(http.StatusOK, resp)
		return
	}
	//存入本地缓存
	utils.AreasCache.Set(conf.HouseAreasCacheIndex, response.Data)
	utils.NewLog().Info("utils.AreasCache.Set:", utils.AreasCache.Len())
	resp["data"] = areas
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
