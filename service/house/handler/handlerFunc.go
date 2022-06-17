package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"ihome/service/house/cache"
	"ihome/service/house/conf"
	"ihome/service/house/kitex_gen"
	"ihome/service/house/model"
	"ihome/service/house/remote"
	po "ihome/service/model"
	user_kitex_gen "ihome/service/user/kitex_gen"
	"ihome/service/utils"
	"strconv"
)

var ctx = context.Background()

func existInMysql(facs []model.FacilityPo, fids []string) bool {
	m := make(map[int]bool, 25)
	for _, v := range facs {
		m[v.Id] = true
	}
	for _, id := range fids {
		index, _ := strconv.ParseInt(id, 10, 32)
		if _, in := m[int(index)]; !in {
			utils.NewLog().Info("fids not in map:", m)
			return false
		}
	}
	return true
}
func PubHouse(sessionId string, houseMap map[string]interface{}) kitex_gen.Response {
	//house存入House表
	facility := houseMap["facility"]
	facilityByte, _ := json.Marshal(facility)
	fids := make([]string, 5)
	json.Unmarshal(facilityByte, &fids)
	//utils.NewLog().Debug("facility:", facility.([]int))
	utils.NewLog().Info("fids:", fids)
	delete(houseMap, "facility")
	//查询redis获取session userId
	sessionResp := model.GetRedis(conf.SessionLoginIndex + "_" + sessionId)
	if utils.RECODE_OK != sessionResp.Errno {
		utils.NewLog().Info("GetRedis error:", sessionResp)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	session := &po.SessionPo{}
	json.Unmarshal(sessionResp.Data, session)
	utils.NewLog().Info("session:", session)
	houseMap["user_id"] = strconv.Itoa(session.ID)
	houseResp := model.SaveMysqlHouse(houseMap)
	if utils.RECODE_OK != houseResp.Errno {
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}

	facByte := GetFacilityIds().Data
	facs := make([]model.FacilityPo, 25)
	json.Unmarshal(facByte, &facs)
	//utils.NewLog().Info("facs:", facs)
	//判断facility是否存在Mysql
	exist := existInMysql(facs, fids)
	if !exist {
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
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
	//直接删除该用户对应的user_houses redis缓存
	model.DeleteKey(conf.UserHouseRedisIndex + "_" + sessionId)
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

//GetUserHouse 使用redis缓存和mysql结合
func GetUserHouse(sessionId string) kitex_gen.Response {
	redisResp := model.GetRedis(conf.UserHouseRedisIndex + "_" + sessionId)
	if string(redisResp.Data) != "" {
		//redis查到
		return utils.HouseResponse(utils.RECODE_OK, redisResp.Data)
	}
	//查不到查数据库
	//先从redis得到user_id
	redisResp = model.GetRedis(conf.SessionLoginIndex + "_" + sessionId)
	if string(redisResp.Data) == "" {
		//redis查不到
		utils.NewLog().Info("session get null:", redisResp)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	session := &po.SessionPo{}
	json.Unmarshal(redisResp.Data, session)
	utils.NewLog().Debug("session:", session)
	mysqlResp := model.GetMysqlHouse(session.ID)
	houses := make([]model.House, 0)
	json.Unmarshal(mysqlResp.Data, &houses)
	utils.NewLog().Debug("houses:", houses)
	//获取areas信息
	areasResp := GetAreas()
	areas := make([]po.AreaPo, 0)
	json.Unmarshal(areasResp.Data, &areas)
	utils.NewLog().Debug("areas:", areas)
	areaMap := make(map[int]string, 0)
	for _, item := range areas {
		areaMap[item.Id] = item.Name
	}
	utils.NewLog().Debug("sorted areas:", areas)
	//获取user信息,调用远程user服务
	userResp := GetUserInfo(sessionId)
	user := po.UserPo{}
	json.Unmarshal(userResp.Data, &user)
	utils.NewLog().Debug("user:", user)
	//housesMap := make(map[string][]po.HousePo)
	houseList := make([]po.HousePo, 0)
	for _, item := range houses {
		housePo := getHousePo(&item, areaMap, &user)
		houseList = append(houseList, housePo)
	}
	//housesMap["houses"] = houseList
	utils.NewLog().Debug("housesList:", houseList)
	//data, _ := json.Marshal(&housesMap)
	data, _ := json.Marshal(&houseList)
	utils.NewLog().Info("redisData:", string(data))
	//更新redis user_house缓存
	saveResp := model.SaveRedis(conf.UserHouseRedisIndex+"_"+sessionId,
		data, conf.UserHouseRedisTimeOut)
	if utils.RECODE_OK != saveResp.Errno {
		utils.NewLog().Error("SaveRedis error:", saveResp)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	return utils.HouseResponse(utils.RECODE_OK, data)
}

func getHousePo(item *model.House, areaMap map[int]string, user *po.UserPo) po.HousePo {
	housePo := po.HousePo{}
	HousesByte, _ := json.Marshal(&item)
	json.Unmarshal(HousesByte, &housePo)
	housePo.Ctime = item.CreatedAt.Format(conf.MysqlTimeFormat)
	housePo.HouseId = item.ID
	housePo.AreaName = areaMap[int(item.AreaId)]
	housePo.ImageUrl = concatImgUrl(conf.NginxUrl, item.Index_image_url)
	housePo.UserAvatar = concatImgUrl(conf.NginxUrl, user.Avatar_url)
	return housePo
}

func GetUserInfo(sessionId string) kitex_gen.Response {
	req := user_kitex_gen.GetUserRequest{SessionId: sessionId}
	res, err := remote.RPC(ctx, conf.UserServiceIndex, req)
	if err != nil {
		utils.NewLog().Info("remote.RPC error:", err)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	response := res.(*user_kitex_gen.Response)
	utils.NewLog().Debug("response:", response)
	return utils.HouseResponse(response.Errno, response.Data)
}

func GetAreas() kitex_gen.Response {
	//先去redis查询缓存信息，查不到再查数据库
	redisAreas := model.GetRedis(conf.RedisAreasIndex).Data
	if "" != string(redisAreas) {
		//redis中存在
		utils.NewLog().Info("GetRedisAreas", string(redisAreas))
		return utils.HouseResponse(utils.RECODE_OK, redisAreas)
	}
	//redis中不存在就查数据库
	response := model.GetMysql(&[]model.Area{})
	areasData := response.Data
	utils.NewLog().Info("model.GetMysqlAreas response:", response.Errmsg)
	if utils.RECODE_OK != response.Errno {
		return response
	}
	//存入redis
	response = model.SaveRedis(conf.RedisAreasIndex, areasData, conf.RedisAreasTimeOut)
	utils.NewLog().Info("model.SaveRedisAreas response:", response)
	if utils.RECODE_OK != response.Errno {
		return response
	}
	response.Data = areasData
	return response
}

func concatImgUrl(nginxDns string, imgUrl string) string {
	return fmt.Sprintf("%s/%s", nginxDns, imgUrl)
}
