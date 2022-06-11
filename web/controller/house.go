package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	house_kitex_gen "ihome/service/house/kitex_gen"
	"ihome/web/conf"
	"ihome/web/model"
	"ihome/web/remote"
	"ihome/web/utils"
	"net/http"
)

//GetAreas 获取地区信息
func GetAreas(ctx *gin.Context) {
	//先走本地缓存
	resp := make(map[string]interface{})
	resp[conf.DataIndex] = ""
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
		resp[conf.DataIndex] = areas
		ctx.JSON(http.StatusOK, resp)
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
