package handler

import (
	"context"
	"encoding/json"
	"ihome/service/elasticsearch"
	"ihome/service/house/cache"
	"ihome/service/house/conf"
	"ihome/service/house/kitex_gen"
	"ihome/service/house/model"
	"ihome/service/house/remote"
	po "ihome/service/model"
	user_kitex_gen "ihome/service/user/kitex_gen"
	"ihome/service/utils"
	"strconv"
	"sync"
)

var ctx = context.Background()

func PubHouse(sessionId string, houseMap map[string]interface{}) kitex_gen.Response {
	//参数中的facilityId
	fids := GetFacilityId(houseMap)
	//查询redis获取session userId
	sessionResp := model.GetRedis(conf.SessionLoginIndex + "_" + sessionId)
	if utils.RECODE_OK != sessionResp.Errno {
		utils.NewLog().Info("GetRedis error:", sessionResp)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	utils.NewLog().Debug("session:", string(sessionResp.Data))
	houseMap["user_id"] = utils.GetFromJson(string(sessionResp.Data), "id")
	houseResp := model.SaveMysqlHouse(houseMap)
	if utils.RECODE_OK != houseResp.Errno {
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	//判断facility是否存在Mysql
	exist := ExistInMysql(fids)
	if !exist {
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	//存house_facility表
	houseFacilities := make([]model.HouseFacilities, 0)
	utils.NewLog().Debug("string(houseResp.Data):", string(houseResp.Data))
	houseId := utils.PareInt(utils.GetFromJson(string(houseResp.Data), "ID"))
	utils.NewLog().Debug("houseID:", houseId)
	for _, id := range fids {
		fac := model.HouseFacilities{HouseId: houseId, FacilityId: id}
		houseFacilities = append(houseFacilities, fac)
	}
	utils.NewLog().Debug("houseFacilities:", houseFacilities)
	houseFacResp := model.SaveMysqlHouseFac(houseFacilities)
	if utils.RECODE_OK != houseResp.Errno {
		utils.NewLog().Info("SaveMysqlHouseFac error:", houseFacResp)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	//启动协程处理
	utils.AntsPool.Pool.Submit(func() {
		//保存house到redis
		houseRedisPo := &po.HouseRedisPo{}
		json.Unmarshal(houseResp.Data, houseRedisPo)
		houseRedisPo.Facilities = fids
		utils.NewLog().Debug("houseRedisPo:", houseRedisPo)
		data, _ := json.Marshal(houseRedisPo)
		model.SaveRedis(utils.ConcatRedisKey(conf.HouseInfoRedisIndex, strconv.Itoa(houseId)),
			data, conf.HouseInfoRedisTimeOut)
		//删除redis user_houses信息
		model.DeleteKey(utils.ConcatRedisKey(conf.UserHousesRedisIndex, sessionId))
		//设置ESHousePO并存入ES
		SetESHouseTask(houseResp.Data)

	})
	return utils.HouseResponse(utils.RECODE_OK, []byte(strconv.Itoa(houseId)))
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
	redisResp := model.GetRedis(utils.ConcatRedisKey(conf.UserHousesRedisIndex, sessionId))
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
	//得到user_id
	userId := utils.PareInt(utils.GetFromJson(string(redisResp.Data), "id"))
	utils.NewLog().Debug("userId:", userId)
	mysqlResp := model.GetMysqlHouses(userId)
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
	//TODO 获取house img信息
	houseList := make([]po.HousePo, 0)
	for _, item := range houses {
		housePo := GetHousePo(&item, areaMap, utils.GetFromJson(string(userResp.Data), "avatar_url"))
		houseList = append(houseList, housePo)
	}
	utils.NewLog().Debug("housesList:", houseList)
	data, _ := json.Marshal(&houseList)
	utils.NewLog().Info("redisData:", string(data))
	//更新redis user_house缓存
	utils.AntsPool.Pool.Submit(func() {
		model.SaveRedis(utils.ConcatRedisKey(conf.UserHousesRedisIndex, sessionId),
			data, conf.UserHousesRedisTimeOut)
	})
	return utils.HouseResponse(utils.RECODE_OK, data)
}

//GetAreas 获取所有areas
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

//UploadHouseImg 上传房子图像
func UploadHouseImg(houseId int, fileType string, imgBase64 string) kitex_gen.Response {
	uploadResp := po.FastDfsClient.UploadImg(imgBase64, fileType)
	if utils.RECODE_OK != uploadResp.Errno {
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	imgPath := string(uploadResp.Data)
	resMap := make(map[string]string)
	resMap["url"] = utils.ConcatImgUrl(conf.NginxUrl, imgPath)
	response := utils.HouseResponse(utils.RECODE_OK, nil)
	res, _ := json.Marshal(&resMap)
	response.Data = res
	SaveHouseImg(houseId, imgPath)
	return response
}

//GetHouseDetail 获取house详细信息
func GetHouseDetail(sessionId string, houseId int) *kitex_gen.HouseDetailResp {
	utils.NewLog().Debug("GetHouseDetail start...")
	house := &po.HouseRedisPo{}
	house.ID = uint(houseId)
	//userInfo
	user := &model.User{}
	//评论
	comments := make([]model.CommentPo, 0)
	//图像url
	imgUrls := make([]string, 0)
	var wg sync.WaitGroup
	wg.Add(3)
	ctx, _ := context.WithTimeout(context.Background(), conf.HouseTaskTimeOut)
	utils.AntsPool.Pool.Submit(GetHouseAndUserTask(&ctx, &wg, sessionId, house, user))
	utils.AntsPool.Pool.Submit(GetCommentsTask(&ctx, &wg, houseId, &comments))
	utils.AntsPool.Pool.Submit(GetHouseImagesTask(&ctx, &wg, houseId, &imgUrls))
	wg.Wait()
	resp := SetHouseDetailResp(house, user, &comments, &imgUrls)
	return resp

}

func GetHouseInfo(houseId int) kitex_gen.Response {
	utils.NewLog().Debug("GetHouseInfo start...")
	//先去redis查
	redisResp := model.GetRedis(utils.ConcatRedisKey(conf.HouseInfoRedisIndex, utils.IntToString(houseId)))
	if string(redisResp.Data) != "" {
		//redis查到直接返回
		utils.NewLog().Debug("model.GetRedis get success....")
		return utils.HouseResponse(utils.RECODE_OK, redisResp.Data)
	}
	//查不到查数据库
	house, err := model.GetMysqlHouseInfo(houseId)
	if err != nil {
		utils.NewLog().Debug("GetMysqlHouseInfo error:", err)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	houseJson, _ := json.Marshal(&house)
	//获取house对应的facilities
	fids := model.GetMysqlHouseFacilityIds(houseId)
	houseRedis := &po.HouseRedisPo{}
	json.Unmarshal(houseJson, houseRedis)
	houseRedis.Facilities = fids
	data, _ := json.Marshal(houseRedis)
	//启动协程更新redis
	utils.AntsPool.Pool.Submit(func() {
		model.SaveRedis(utils.ConcatRedisKey(conf.HouseInfoRedisIndex, utils.IntToString(houseId)),
			data, conf.HouseInfoRedisTimeOut)
	})
	return utils.HouseResponse(utils.RECODE_OK, data)
}

//GetUserInfo 获取用户信息
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

//GetHouseImgUrls 获得house的图像
func GetHouseImgUrls(houseId int) kitex_gen.Response {
	//先查redis
	utils.NewLog().Debug("GetHouseImgUrls start...")
	redisResp := model.GetRedis(utils.ConcatRedisKey(conf.HouseImgRedisIndex, houseId))
	if string(redisResp.Data) != "" {
		//查到直接返回
		return redisResp
	}
	//查不到去查数据库
	imgUrls := model.GetMysqlHouseImg(houseId)
	data, _ := json.Marshal(&imgUrls)
	//更新redis
	utils.AntsPool.Pool.Submit(func() {
		model.SaveRedis(utils.ConcatRedisKey(conf.HouseImgRedisIndex, houseId),
			data, conf.HouseImgRedisTimeOut)
	})
	return utils.HouseResponse(utils.RECODE_OK, data)
}

//SearchHouse 按条件搜索房源
func SearchHouse(req *kitex_gen.HouseSearchReq) *kitex_gen.HouseSearchResp {
	searchReq := &elasticsearch.HouseSearchReq{}
	searchResp := &kitex_gen.HouseSearchResp{}
	//searchReq := &elasticsearch.HouseSearchReq{}
	SetHouseSearchReq(searchReq, req)
	filter := searchReq.ToFilter()
	ctx, _ := context.WithTimeout(context.Background(), conf.ESTaskTimeOut)
	result, err := elasticsearch.HouseES.Search(ctx, filter)
	if err != nil {
		utils.NewLog().Info("elasticsearch.HouseES.Search error:", err)
		searchResp.Errno = utils.RECODE_SERVERERR
		searchResp.Errmsg = utils.RecodeText(utils.RECODE_SERVERERR)
		return searchResp
	}
	searchResp.Errno = utils.RECODE_OK
	searchResp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	//设置搜索结果
	searchResp.Data = SetHouseSearchData(req, filter, result)
	return searchResp
}

//GetHouseHomeIndex 主页房源轮播
func GetHouseHomeIndex(sessionId string) *kitex_gen.HouseSearchResp {
	indexResp := &kitex_gen.HouseSearchResp{}
	utils.NewLog().Debug("GetHouseHomeIndex sessionId:", sessionId)
	redisResult := model.GetRedis(utils.ConcatRedisKey(conf.HouseHomePageRedisIndex, sessionId))
	if string(redisResult.Data) != "" {
		//不为空直接返回
		indexResp.Errno = utils.RECODE_OK
		indexResp.Errmsg = utils.RecodeText(utils.RECODE_OK)
		houses := make([]*kitex_gen.HouseSearchInfo, 0)
		json.Unmarshal(redisResult.Data, &houses)
		indexResp.Data = &kitex_gen.SearchResp{Houses: houses}
		return indexResp
	}
	//查不到查ES取前十的结果
	searchReq := &kitex_gen.HouseSearchReq{Page: 1, Size: 10}
	searchReq.SessionId = sessionId
	searchResp := SearchHouse(searchReq)
	indexResp.Errno = utils.RECODE_OK
	indexResp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	indexResp.Data = &kitex_gen.SearchResp{Houses: searchResp.Data.Houses}
	return indexResp
}
