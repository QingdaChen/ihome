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
func GetHousesInfo(ctx *gin.Context) {
	houses := model.HouseVO{Houses: []model.House{}}
	ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_OK, houses))
}

//PubHouses Post 发布房子信息
func PubHouses(ctx *gin.Context) {
	utils.NewLog().Info("PubHouses start")
	params, _ := ioutil.ReadAll(ctx.Request.Body)
	utils.NewLog().Debug("params:", params)
	req := house_kitex_gen.PubHouseRequest{Params: params}
	res, err2 := remote.RPC(ctx, conf.HouseServiceIndex, req)
	if err2 != nil {
		utils.NewLog().Info("remote.RPC error:", err2)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return
	}
	response := res.(*house_kitex_gen.Response)
	ctx.JSON(http.StatusOK, utils.Response(response.Errno, response.Data))
	return
}
