package handler

import (
	"context"
	"encoding/json"
	"ihome/service/elasticsearch"
	"ihome/service/house/conf"
	"ihome/service/house/kitex_gen"
	"ihome/service/house/model"
	po "ihome/service/model"
	"ihome/service/utils"
)

func ExistInMysql(fids []int) bool {
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
func GetFacilityId(houseMap map[string]interface{}) []int {
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
func GetHousePo(item *model.House, areaMap map[int]string, avatarUrl string) po.HousePo {
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
func SaveHouseImg(houseId int, imgUrl string) {
	//提交协程任务，将imgPath存入数据库
	utils.AntsPool.Pool.Submit(func() {
		//先保存houseId 和 imgUrl
		model.SaveMysqlHouseIdImg(houseId, imgUrl)
		imgUrls := model.GetMysqlHouseImg(houseId)
		utils.NewLog().Debug("imgUrls:", imgUrls)
		data, _ := json.Marshal(&imgUrls)
		//保存redis
		model.SaveRedis(utils.ConcatRedisKey(conf.HouseImgRedisIndex, utils.IntToString(houseId)),
			data, conf.HouseImgRedisTimeOut)
		//修改ES
		if len(imgUrls) == 1 {
			//housePo := &elasticsearch.ESHousePo{HouseId: int64(houseId), ImageUrl: imgUrl}
			m := make(map[string]interface{}, 0)
			m["house_id"] = houseId
			m["image_url"] = imgUrl
			//marshal, _ := json.Marshal(m)
			//utils.NewLog().Debug("housePo:", string(marshal))
			housePos := []map[string]interface{}{m}
			ctx, _ := context.WithTimeout(context.Background(), conf.ESTaskTimeOut)
			err := elasticsearch.HouseES.BatchUpdate(ctx, housePos)
			if err != nil {
				utils.NewLog().Info("elasticsearch.HouseES.BatchUpdate error:", err)
			}
		}
	})
}
func SetHouseDetailResp(house *po.HouseRedisPo, user *model.User,
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
	houseDetail.Hid = int64(house.ID)
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

func SetHouseSearchReq(searchReq *elasticsearch.HouseSearchReq, req *kitex_gen.HouseSearchReq) {
	//searchReq.Page = int(req.Page)
	//searchReq.AreaId = req.AreaId
	//searchReq.MinPrice = int(req.MinPrice)
	//searchReq.MaxPrice = int(req.MaxPrice)
	reqJson, _ := json.Marshal(req)
	json.Unmarshal(reqJson, searchReq)
	searchJson, _ := json.Marshal(searchReq)
	utils.NewLog().Debug("searchReq:", string(searchJson))
}

func SetHouseSearchData(req *kitex_gen.HouseSearchReq,
	filter *elasticsearch.ESSearch, result []*elasticsearch.ESHousePo) *kitex_gen.SearchResp {
	searchData := &kitex_gen.SearchResp{}
	searchData.CurrentPage = req.Page
	searchData.TotalPage = int32((filter.Size / int(req.Size)) + 1)
	searchData.Houses = make([]*kitex_gen.HouseSearchInfo, 0)
	for _, house := range result {
		info := &kitex_gen.HouseSearchInfo{}
		data, _ := json.Marshal(house)
		utils.NewLog().Debug("data:", string(data))
		json.Unmarshal(data, info)
		infoJson, _ := json.Marshal(info)
		utils.NewLog().Debug("info:", string(infoJson))
		info.ImgUrl = utils.ConcatImgUrl(conf.NginxUrl, house.ImageUrl)
		info.UserAvatar = utils.ConcatImgUrl(conf.NginxUrl, house.UserAvatar)
		info.AreaName = house.AreaName
		searchData.Houses = append(searchData.Houses, info)
	}
	//启动协程存redis
	utils.AntsPool.Pool.Submit(func() {
		data, _ := json.Marshal(searchData)
		utils.NewLog().Debug("SetHouseSearchData sessionId:", req.SessionId)
		model.SaveRedis(utils.ConcatRedisKey(conf.HouseHomePageRedisIndex, req.SessionId),
			data, conf.HouseHomePageRedisTimeOut)
	})
	return searchData
}
