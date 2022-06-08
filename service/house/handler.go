package main

import (
	"context"
	"ihome/service/house/kitex_gen"
	"ihome/service/house/model"
	"ihome/service/utils"
)

// HouseServiceImpl implements the last service interface defined in the IDL.
type HouseServiceImpl struct{}

// GetArea implements the HouseServiceImpl interface.
func (s *HouseServiceImpl) GetArea(ctx context.Context, req *kitex_gen.AreaRequest) (resp *kitex_gen.Response, err error) {
	//先去redis查询缓存信息，查不到再查数据库
	redisAreas := model.GetRedisAreas().Data
	if "" != string(redisAreas) {
		//redis中存在
		utils.NewLog().Info("GetRedisAreas", string(redisAreas))
		response := utils.HouseResponse(utils.RECODE_OK, redisAreas)
		return &response, nil
	}
	//redis中不存在就查数据库
	response := model.GetMysqlAreas()
	areasData := response.Data
	utils.NewLog().Info("model.GetMysqlAreas response:", response.Errmsg)
	if utils.RECODE_OK != response.Errno {
		return &response, nil
	}
	//存入redis
	response = model.SaveRedisAreas(areasData)
	utils.NewLog().Info("model.SaveRedisAreas response:", response)
	if utils.RECODE_OK != response.Errno {
		return &response, nil
	}
	response.Data = areasData
	return &response, nil
}
