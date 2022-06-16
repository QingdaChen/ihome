package handler

import (
	"encoding/json"
	"ihome/service/house/cache"
	"ihome/service/house/conf"
	"ihome/service/house/kitex_gen"
	"ihome/service/house/model"
	"ihome/service/utils"
	"strconv"
)

func PubHouse(houseMap map[string]interface{}) kitex_gen.Response {
	//house存入House表
	facility := houseMap["facility"]
	facilityByte, _ := json.Marshal(facility)
	fids := make([]string, 5)
	json.Unmarshal(facilityByte, &fids)
	//utils.NewLog().Debug("facility:", facility.([]int))
	utils.NewLog().Info("fids:", fids)
	delete(houseMap, "facility")
	houseResp := model.SaveMysqlHouse(houseMap)
	if utils.RECODE_OK != houseResp.Errno {
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}

	facByte := GetFacilityIds().Data
	facs := make([]model.FacilityPo, 25)
	json.Unmarshal(facByte, &facs)
	//utils.NewLog().Info("facs:", facs)
	//判断facility是否存在Mysql
	m := make(map[int]bool, 25)
	for _, v := range facs {
		m[v.Id] = true
	}
	for _, id := range fids {
		index, _ := strconv.ParseInt(id, 10, 32)
		if _, in := m[int(index)]; !in {
			utils.NewLog().Info("fids not in map:", m)
			return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
		}

	}
	//存house_facility表
	HouseFacs := make([]model.HouseFacilities, 0)
	houseId, _ := strconv.ParseInt(string(houseResp.Data), 10, 32)
	for _, idStr := range fids {
		id, _ := strconv.ParseInt(idStr, 10, 32)
		fac := model.HouseFacilities{HouseId: int(houseId), FacilityId: int(id)}
		HouseFacs = append(HouseFacs, fac)
	}
	utils.NewLog().Debug("HouseFacs:", HouseFacs)
	houseFacResp := model.SaveMysqlHouseFac(HouseFacs)
	if utils.RECODE_OK != houseResp.Errno {
		utils.NewLog().Info("SaveMysqlHouseFac error:", houseFacResp)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	data, _ := json.Marshal(&HouseFacs[0])
	utils.NewLog().Info("facilities data:", data)
	return utils.HouseResponse(utils.RECODE_OK, data)
}

func GetFacilityIds() kitex_gen.Response {
	//先查本地缓存
	data, err := cache.FacilityCache.Get(conf.FacilityIndex)
	if err == nil {
		response := utils.HouseResponse(utils.RECODE_OK, data)
		return response
	}
	//再查redis
	redisResp := model.GetRedis(conf.FacilityIndex)
	if string(redisResp.Data) != "" {
		utils.NewLog().Info("redisResp:", string(redisResp.Errmsg))
		//redis查到了
		//更新本地缓存
		cache.FacilityCache.Set(conf.FacilityIndex, redisResp.Data)
		response := utils.HouseResponse(utils.RECODE_OK, redisResp.Data)
		return response
	}
	//最后查Mysql
	mysqlResp := model.GetMysql(&[]model.Facility{})
	if utils.RECODE_OK != utils.RECODE_OK {
		//系统错误
		utils.NewLog().Error("GetMysql error:", mysqlResp)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	//保存redis
	model.SaveRedis(conf.FacilityIndex, mysqlResp.Data, conf.FacilityRedisTimeOut)
	//更新本地缓存
	cache.FacilityCache.Set(conf.FacilityIndex, mysqlResp.Data)
	return utils.HouseResponse(utils.RECODE_OK, mysqlResp.Data)
}
