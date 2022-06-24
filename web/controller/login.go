package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"ihome/remote"
	user_kitex_gen "ihome/service/user/kitex_gen"
	"ihome/web/conf"
	"ihome/web/model"
	"ihome/web/utils"
	"net/http"
)

var ctx, _ = context.WithTimeout(context.Background(), conf.RPCTimeOut)

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

	utils.NewLog().Info("GetSMSCd..." + phone + ":" + imgCode + ":" + uuid)
	//发送短信
	req := user_kitex_gen.SMSRequest{Phone: phone, ImgCode: imgCode, Uuid: uuid}
	res, err := remote.RPC(ctx, conf.UserServiceIndex, req)
	if err != nil {
		utils.NewLog().Info("SendSMS...", err)
	}
	ctx.JSON(http.StatusOK, res)

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
	req := user_kitex_gen.RegRequest{Phone: register.Phone,
		Password: register.Password, SmsCode: register.SmsCode}
	res, err2 := remote.RPC(ctx, conf.UserServiceIndex, req)
	if err2 != nil {
		utils.NewLog().Info("rpc Register error:", err)
	}
	ctx.JSON(http.StatusOK, res)

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

	req := user_kitex_gen.LoginRequest{Phone: loginUser.Phone, Password: loginUser.Password}
	res, err2 := remote.RPC(ctx, conf.UserServiceIndex, req)
	if err2 != nil {
		utils.NewLog().Error("rpc LoginRequest error:", err2)

		return
	}
	response := res.(*user_kitex_gen.Response)
	//登录成功就保存session
	utils.NewLog().Info("login response:", string(response.Data))
	ctx.SetCookie(conf.LoginCookieName, string(response.Data),
		conf.LoginCookieTimeOut, "/", "", false, true)
	ctx.JSON(http.StatusOK, res)
	return
}

//SessionAuth session鉴权
func SessionAuth(ctx *gin.Context) {
	sessionId, err := ctx.Cookie(conf.LoginCookieName) // 获得cookie seesionId
	utils.NewLog().Info("sessionId:", sessionId)
	if err != nil || sessionId == "" {
		//cookie未存在或过期直接返回
		utils.NewLog().Info("cookie未存在或过期直接返回:")
		ctx.Abort()
		ctx.Redirect(http.StatusTemporaryRedirect, conf.LoginHtmlLocation+"/?resp=未登录") //307 临时重定向
		return
	}

	//连接user服务 查询session
	req := user_kitex_gen.SessionAuthRequest{SessionId: sessionId}
	res, err2 := remote.RPC(context.Background(), conf.UserServiceIndex, req)
	utils.NewLog().Info("SessionAuthRequest result:", res, err2)
	if err2 != nil {
		utils.NewLog().Error("rpc SessionAuth error:", err2)
		ctx.Abort()
		ctx.Redirect(http.StatusTemporaryRedirect, conf.LoginHtmlLocation+"/?resp=未登录") //307 临时重定向
		return
	}

	response := res.(*user_kitex_gen.Response)
	utils.NewLog().Info("rpc response:", response)
	//验证失败
	if utils.RECODE_OK != response.Errno {
		utils.NewLog().Info("SessionAuth fail:", err2, response)
		ctx.Abort()
		//ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SESSIONERR, nil))
		//重定向到login.html
		ctx.Redirect(http.StatusTemporaryRedirect, conf.LoginHtmlLocation+"/?resp=未登录") //307 临时重定向

		return
	}
	ctx.Next()
	return
}

func Index(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, "http://192.168.31.219:8088/home")
}
