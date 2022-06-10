package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	house_kitex_gen "ihome/service/house/kitex_gen"
	"ihome/service/house/kitex_gen/houseservice"
	"ihome/web/conf"
	"ihome/web/model"
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
	var result interface{}
	result, resp = utils.GetService(conf.HouseServiceIndex)
	if utils.RECODE_OK != resp[conf.ErrorNoIndex].(string) {
		ctx.JSON(http.StatusOK, resp)
		return
	}
	service := result.(houseservice.Client)
	req := &house_kitex_gen.AreaRequest{}
	response, _ := service.GetArea(ctx, req)
	utils.Resp(resp, response.Errno)
	if utils.RECODE_OK != response.Errno {
		resp["data"] = ""
		ctx.JSON(http.StatusOK, resp)
		return
	}
	//反序列化
	err := json.Unmarshal(response.Data, &areas)
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
