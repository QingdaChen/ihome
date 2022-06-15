package controller

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	user_kitex_gen "ihome/service/user/kitex_gen"
	"ihome/web/conf"
	"ihome/web/model"
	"ihome/web/remote"
	"ihome/web/utils"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

//GetSession 获取session信息
func GetSession(ctx *gin.Context) {

	utils.NewLog().Info("GetSession start...")
	sessionId, err := ctx.Cookie(conf.LoginCookieName)
	if err != nil || sessionId == "" {
		//sessionId 不存在或者过期直接返回
		utils.NewLog().Info("ctx.Cookie error:", err)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SESSIONERR, nil))
		return
	}
	req := user_kitex_gen.SessionRequest{SessionId: sessionId}
	res, err2 := remote.RPC(ctx, conf.UserServiceIndex, req)
	if err2 != nil {
		utils.NewLog().Info("remote.RPC error:", err)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return
	}
	var user model.SessionVo
	response := res.(*user_kitex_gen.Response)
	//未登录...
	if utils.RECODE_OK != response.Errno {
		utils.NewLog().Info("session nil:", response.Errmsg)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return
	}
	err = json.Unmarshal(response.Data, &user)
	if err != nil {
		utils.NewLog().Error("json.Unmarshal error:", err)
		ctx.JSON(http.StatusOK, response)
	}
	utils.NewLog().Info("response:", user)
	ctx.JSON(http.StatusOK, utils.Response(response.Errno, user))

}

//DeleteSession 删除session信息,退出登录
func DeleteSession(ctx *gin.Context) {

	utils.NewLog().Info("DeleteSession start...")
	sessionId, err := ctx.Cookie(conf.LoginCookieName)
	if err != nil || sessionId == "" {
		//sessionId 不存在或者过期直接返回
		utils.NewLog().Info("ctx.Cookie error:", err)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SESSIONERR, nil))
		return
	}
	req := user_kitex_gen.SessionDeleteRequest{SessionId: sessionId}
	res, err2 := remote.RPC(ctx, conf.UserServiceIndex, req)
	if err2 != nil {
		utils.NewLog().Info("remote.RPC error:", err)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return
	}
	response := res.(*user_kitex_gen.Response)
	utils.NewLog().Info("response:", response)
	ctx.JSON(http.StatusOK, utils.Response(response.Errno, nil))

}

//GetUserInfo 获取用户信息
func GetUserInfo(ctx *gin.Context) {
	//TODO 限流
	utils.NewLog().Info("GetUserInfo start...")
	sessionId, err := ctx.Cookie(conf.LoginCookieName)
	if err != nil || sessionId == "" {
		//sessionId 不存在或者过期直接返回
		utils.NewLog().Info("ctx.Cookie error:", err)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SESSIONERR, nil))
		return
	}
	req := user_kitex_gen.GetUserRequest{SessionId: sessionId}
	res, err2 := remote.RPC(ctx, conf.UserServiceIndex, req)
	if err2 != nil {
		utils.NewLog().Info("remote.RPC error:", err)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return
	}
	var user model.UserVo
	response := res.(*user_kitex_gen.Response)
	//反序列化
	err = json.Unmarshal(response.Data, &user)
	if err != nil {
		utils.NewLog().Error("json.Unmarshal error:", err)
		ctx.JSON(http.StatusOK, response)
	}
	utils.NewLog().Info("response:", user)
	user.Avatar_url = fmt.Sprintf("%s/%s", conf.NginxUrl, user.Avatar_url)
	ctx.JSON(http.StatusOK, utils.Response(response.Errno, user))
}

//UpdateUserInfo 更新用户信息
func UpdateUserInfo(ctx *gin.Context) {
	//TODO 限流
	utils.NewLog().Debug("UpdateUserInfo start...")
	updateUserMap := make(map[string]string, 5)
	params, _ := ioutil.ReadAll(ctx.Request.Body)
	json.Unmarshal(params, &updateUserMap)
	utils.NewLog().Println("updateUserMap:", updateUserMap)

	sessionId, err := ctx.Cookie(conf.LoginCookieName)
	if err != nil || sessionId == "" {
		//sessionId 不存在或者过期直接返回
		utils.NewLog().Info("ctx.Cookie error:", err)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SESSIONERR, nil))
		return
	}
	req := user_kitex_gen.UpdateUserRequest{SessionId: sessionId, Data: updateUserMap}
	res, err2 := remote.RPC(ctx, conf.UserServiceIndex, req)
	if err2 != nil {
		utils.NewLog().Info("remote.RPC error:", err)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return
	}
	//更新成功
	response := res.(*user_kitex_gen.Response)
	utils.NewLog().Info("response:", response)
	ctx.JSON(http.StatusOK, utils.Response(response.Errno, nil))
}

//PostAvatar 上传头像
func PostAvatar(ctx *gin.Context) {
	utils.NewLog().Info("PostAvatar start...")
	//获取cookie
	sessionId, err := ctx.Cookie(conf.LoginCookieName)
	if err != nil || sessionId == "" {
		//sessionId 不存在或者过期直接返回
		utils.NewLog().Info("ctx.Cookie error:", err)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SESSIONERR, nil))
		return
	}
	file, head, err2 := ctx.Request.FormFile("avatar")
	fileType := strings.Split(head.Filename, ".")[1]
	utils.NewLog().Debug("fileType: ", fileType)
	if err2 != nil {
		utils.NewLog().Info("FormFile error:", err)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return
	}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)
	//base64编码
	imgBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	//utils.NewLog().Info("imgBase64:  ", imgBase64)
	//调用图像上传服务
	req := user_kitex_gen.UploadImgRequest{SessionId: sessionId, FileType: fileType, ImgBase64: imgBase64}
	res, err2 := remote.RPC(ctx, conf.UserServiceIndex, req)
	if err2 != nil {
		utils.NewLog().Info("remote.RPC error:", err2)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return
	}
	//更新成功
	response := res.(*user_kitex_gen.Response)
	utils.NewLog().Info("res:", string(response.Data))
	userVo := &model.UserVo{}
	userVo.Avatar_url = fmt.Sprintf("%s/%s", conf.NginxUrl, string(response.Data))
	utils.NewLog().Info("response:", userVo.Avatar_url)
	ctx.JSON(http.StatusOK, utils.Response(response.Errno, userVo))

}

func Test(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, "http://192.168.31.219:8088/home/login.html")
	return
}
