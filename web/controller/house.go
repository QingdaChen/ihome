package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	house_kitex_gen "ihome/service/house/kitex_gen"
	"ihome/web/conf"
	"ihome/web/model"
	"ihome/web/remote"
	"ihome/web/utils"
	"io/ioutil"
	"net/http"
)

//GetAreas 获取地区信息
func GetAreas(ctx *gin.Context) {
	//先走本地缓存
	//缓存中查
	cacheAreas, _ := utils.AreasCache.Get(conf.HouseAreasCacheIndex)
	var areas []model.Area
	if cacheAreas != nil {
		utils.NewLog().Info("cacheAreas:", string(cacheAreas))
		err := json.Unmarshal(cacheAreas, &areas)
		if err != nil {
			utils.NewLog().Error("json.Unmarshal error:", err)
			ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
			return
		}
		utils.NewLog().Info("cache areas:", areas)
		//直接返回
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_OK, areas))
		return
	}
	//查不到远程请求house服务查询

	req := house_kitex_gen.AreaRequest{}
	res, err := remote.RPC(ctx, conf.HouseServiceIndex, req)
	if err != nil {
		utils.NewLog().Error("rpc GetAreas error")
		return
	}
	//获得返回数据
	response := res.(*house_kitex_gen.Response)
	//utils.NewLog().Info("response.Data:", response.Data)
	//反序列化
	err2 := json.Unmarshal(response.Data, &areas)
	if err2 != nil {
		utils.NewLog().Info("json.Unmarshal(response.Data, &areas) error", err)
		ctx.JSON(http.StatusOK, res)
		return
	}
	utils.NewLog().Info("areas:", areas)
	ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_OK, areas))
	//存入本地缓存
	utils.AreasCache.Set(conf.HouseAreasCacheIndex, response.Data)
	utils.NewLog().Info("utils.AreasCache.Set:", utils.AreasCache.Len())
	return

}

//GetHousesInfo 获取房子信息
func GetUserHouses(ctx *gin.Context) {
	utils.NewLog().Debug("PubHouses start")
	//获得cookie
	sessionId, err := ctx.Cookie(conf.LoginCookieName)
	if err != nil || sessionId == "" {
		//sessionId 不存在或者过期直接返回
		utils.NewLog().Info("ctx.Cookie error:", err)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SESSIONERR, nil))
		return
	}
	req := house_kitex_gen.GetUserHouseRequest{SessionId: sessionId}
	res, err2 := remote.RPC(ctx, conf.HouseServiceIndex, req)
	if err2 != nil {
		utils.NewLog().Info("remote.RPC error:", err2)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return
	}
	response := res.(*house_kitex_gen.Response)
	resp := make(map[string][]model.UserHouseVo)
	userHouses := make([]model.UserHouseVo, 0)
	json.Unmarshal(response.Data, &userHouses)
	resp["houses"] = userHouses
	ctx.JSON(http.StatusOK, utils.Response(response.Errno, resp))
	return
}

//PubHouses Post 发布房子信息
func PubHouses(ctx *gin.Context) {
	utils.NewLog().Info("PubHouses start")
	//获得cookie
	sessionId, err := ctx.Cookie(conf.LoginCookieName)
	if err != nil || sessionId == "" {
		//sessionId 不存在或者过期直接返回
		utils.NewLog().Info("ctx.Cookie error:", err)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SESSIONERR, nil))
		return
	}
	params, _ := ioutil.ReadAll(ctx.Request.Body)
	//	utils.NewLog().Debug("params:", params)
	req := house_kitex_gen.PubHouseRequest{SessionId: sessionId, Params: params}
	res, err2 := remote.RPC(ctx, conf.HouseServiceIndex, req)
	if err2 != nil {
		utils.NewLog().Info("remote.RPC error:", err2)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return
	}
	response := res.(*house_kitex_gen.Response)
	resp := utils.Response(response.Errno, nil)
	m := make(map[string]string, 0)
	m["house_id"] = string(response.Data)
	resp["data"] = m
	ctx.JSON(http.StatusOK, resp)
	return
}

//UploadHouseImg 上传house图像
func UploadHouseImg(ctx *gin.Context) {
	utils.NewLog().Info("UploadHouseImg start...")
	houseId := utils.PareInt(ctx.Param("id"))
	utils.NewLog().Debug("houseId:", houseId)
	////TODO 上传图像校验
	resp, fileType, imgBase64 := utils.CheckFile(ctx, "house_image")
	if utils.RECODE_OK != utils.RECODE_OK {
		ctx.JSON(http.StatusOK, resp)
		return
	}
	//调用house的图像上传服务
	req := house_kitex_gen.UploadHouseImgReq{HouseId: int64(houseId), FileType: fileType, ImgBase64: imgBase64}
	res, err2 := remote.RPC(ctx, conf.HouseServiceIndex, req)
	if err2 != nil {
		utils.NewLog().Info("remote.RPC error:", err2)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return
	}
	//更新成功
	response := res.(*house_kitex_gen.Response)
	utils.NewLog().Info("res:", string(response.Data))
	type Result struct {
		Url string `json:"url"`
	}
	result := &Result{}
	json.Unmarshal(response.Data, &result)
	ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_OK, result))
}

func GetHouseDetail(ctx *gin.Context) {
	utils.NewLog().Debug("GetHouseDetail start...")
	//获得cookie
	sessionId, err := ctx.Cookie(conf.LoginCookieName)
	if err != nil || sessionId == "" {
		//sessionId 不存在或者过期直接返回
		utils.NewLog().Info("ctx.Cookie error:", err)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SESSIONERR, nil))
		return
	}
	//获得houseId
	houseId := utils.PareInt(ctx.Param("id"))
	//调用house服务获得houseDetail
	req := house_kitex_gen.GetHouseDetailReg{SessionId: sessionId, HouseId: int64(houseId)}
	res, err2 := remote.RPC(ctx, conf.HouseServiceIndex, req)
	if err2 != nil {
		utils.NewLog().Info("remote.RPC error:", err2)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return
	}
	//更新成功
	response := res.(*house_kitex_gen.HouseDetailResp)
	utils.NewLog().Info("res:", response.Data)
	ctx.JSON(http.StatusOK, response)
}
