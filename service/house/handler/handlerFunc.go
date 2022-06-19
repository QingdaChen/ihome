package handler

import (
	"context"
	"encoding/json"
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
	fids := getFacilityId(houseMap)
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
	exist := existInMysql(fids)
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
	//data, _ := json.Marshal(&houseFacilities[0])
	//utils.NewLog().Info("facilities data:", data)
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
		housePo := getHousePo(&item, areaMap, utils.GetFromJson(string(userResp.Data), "avatar_url"))
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
	saveHouseImg(houseId, imgPath)
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
	resp := setHouseDetailResp(house, user, &comments, &imgUrls)
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

//
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

func existInMysql(fids []int) bool {
	facByte := GetFacilityIds().Data
	facs := make([]model.FacilityPo, 25)
	json.Unmarshal(facByte, &facs)
	m := make(map[int]bool, 25)
	for _, v := range facs {
		m[v.Id] = true
	}
	for _, id := range fids {
		if _, in := m[id]; !in {
			utils.NewLog().Info("fids not in map:", m)
			return false
		}
	}
	return true
}
func getFacilityId(houseMap map[string]interface{}) []int {
	facility := houseMap["facility"]
	utils.NewLog().Debug("facility:", facility)
	facilityByte, _ := json.Marshal(facility)
	fids := make([]string, 5)
	json.Unmarshal(facilityByte, &fids)
	utils.NewLog().Info("fids:", fids)
	delete(houseMap, "facility")
	res := make([]int, 0)
	for _, idStr := range fids {
		res = append(res, utils.PareInt(idStr))
	}
	return res
}
func getHousePo(item *model.House, areaMap map[int]string, avatarUrl string) po.HousePo {
	housePo := po.HousePo{}
	HousesByte, _ := json.Marshal(&item)
	json.Unmarshal(HousesByte, &housePo)
	housePo.Ctime = item.CreatedAt.Format(conf.MysqlTimeFormat)
	housePo.HouseId = item.ID
	housePo.AreaName = areaMap[int(item.AreaId)]
	housePo.ImageUrl = utils.ConcatImgUrl(conf.NginxUrl, item.Index_image_url)
	housePo.UserAvatar = utils.ConcatImgUrl(conf.NginxUrl, avatarUrl)
	return housePo
}
func saveHouseImg(houseId int, imgUrl string) {
	//提交协程任务，将imgPath存入数据库
	utils.AntsPool.Pool.Submit(func() {
		//先保存houseId 和 imgUrl
		model.SaveMysqlHouseIdImg(houseId, imgUrl)
		imgUrls := model.GetMysqlHouseImg(houseId)
		utils.NewLog().Debug("imgUrls:", imgUrls)
		data, _ := json.Marshal(&imgUrls)
		model.SaveRedis(utils.ConcatRedisKey(conf.HouseImgRedisIndex, utils.IntToString(houseId)),
			data, conf.HouseImgRedisTimeOut)
	})
}
func setHouseDetailResp(house *po.HouseRedisPo, user *model.User,
	comments *[]model.CommentPo, imgUrls *[]string) *kitex_gen.HouseDetailResp {
	resp := &kitex_gen.HouseDetailResp{}
	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	data := kitex_gen.HouseDetailData{}
	data.UserId = int64(house.UserId)
	houseDetail := &kitex_gen.HouseDetail{}
	//set house信息
	houseJson, _ := json.Marshal(house)
	json.Unmarshal(houseJson, houseDetail)
	//set user 信息
	utils.NewLog().Debug("user:", user)
	houseDetail.UserAvatar = utils.ConcatImgUrl(conf.NginxUrl, user.Avatar_url)
	houseDetail.UserId = int64(user.ID)
	houseDetail.UserName = user.Name
	utils.NewLog().Debug("houseDetail:", houseDetail)
	houseDetail.ImgUrls = *imgUrls
	//评论信息
	coms := make([]*kitex_gen.CommentInfo, 0)
	for _, item := range *comments {
		m := &kitex_gen.CommentInfo{}
		m.Comment = item.Comment
		m.Ctime = item.Ctime
		m.UserName = item.UserName
		coms = append(coms, m)
	}
	houseDetail.Comments = coms
	utils.NewLog().Debug("houseDetail:", houseDetail)
	data.House = houseDetail
	resp.Data = &data
	utils.NewLog().Debug("resp:", resp)
	return resp

}
