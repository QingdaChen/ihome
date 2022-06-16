package main

import (
	"context"
	"encoding/json"
	"ihome/service/house/conf"
	"ihome/service/house/handler"
	"ihome/service/house/kitex_gen"
	"ihome/service/house/model"
	"ihome/service/utils"
)

// HouseServiceImpl implements the last service interface defined in the IDL.
type HouseServiceImpl struct{}

// GetArea implements the HouseServiceImpl interface.
func (s *HouseServiceImpl) GetArea(ctx context.Context, req *kitex_gen.AreaRequest) (resp *kitex_gen.Response, err error) {
	//先去redis查询缓存信息，查不到再查数据库
	redisAreas := model.GetRedis(conf.RedisAreasIndex).Data
	if "" != string(redisAreas) {
		//redis中存在
		utils.NewLog().Info("GetRedisAreas", string(redisAreas))
		response := utils.HouseResponse(utils.RECODE_OK, redisAreas)
		return &response, nil
	}
	//redis中不存在就查数据库
	response := model.GetMysql(&[]model.Area{})
	areasData := response.Data
	utils.NewLog().Info("model.GetMysqlAreas response:", response.Errmsg)
	if utils.RECODE_OK != response.Errno {
		return &response, nil
	}
	//存入redis
	response = model.SaveRedis(conf.RedisAreasIndex, areasData, conf.RedisAreasTimeOut)
	utils.NewLog().Info("model.SaveRedisAreas response:", response)
	if utils.RECODE_OK != response.Errno {
		return &response, nil
	}
	response.Data = areasData
	return &response, nil
}

// PubHouse implements the HouseServiceImpl interface.
func (s *HouseServiceImpl) PubHouse(ctx context.Context, req *kitex_gen.PubHouseRequest) (resp *kitex_gen.Response, err error) {
	utils.NewLog().Debug("PubHouse start...")
	houseMap := make(map[string]interface{}, 10)
	err = json.Unmarshal(req.Params, &houseMap)
	if err != nil {
		utils.NewLog().Info("json.Unmarshal error:", err)
		response := utils.HouseResponse(utils.RECODE_SERVERERR, nil)
		return &response, nil
	}
	pubResp := handler.PubHouse(houseMap)
	if utils.RECODE_OK != pubResp.Errno {
		utils.NewLog().Info("PubHouse error:", pubResp)
	}
	return &pubResp, nil

}
