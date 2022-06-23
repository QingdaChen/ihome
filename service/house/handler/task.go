package handler

import (
	"context"
	"encoding/json"
	"ihome/service/elasticsearch"
	"ihome/service/house/conf"
	"ihome/service/house/model"
	po "ihome/service/model"
	"ihome/service/utils"
	"sync"
)

func GetHouseAndUserTask(ctx *context.Context, wg *sync.WaitGroup, sessionId string,
	house *po.HouseRedisPo, user *model.User) func() {
	return func() {
		defer wg.Done()
		select {
		case <-(*ctx).Done():
			return
		default:
			//获取houseInfo
			houseResp := GetHouseInfo(int(house.ID))
			if utils.RECODE_OK != houseResp.Errno {
				utils.NewLog().Debug("GetHouseInfo error:", houseResp)
				return
			}
			//获取userInfo
			userResp := GetUserInfo(sessionId)
			if utils.RECODE_OK != houseResp.Errno {
				utils.NewLog().Debug("getUserInfo error:", userResp)
				return
			}
			json.Unmarshal(houseResp.Data, house)
			json.Unmarshal(userResp.Data, user)
			return
		}

	}
}

func GetCommentsTask(ctx *context.Context, wg *sync.WaitGroup, houseId int, comments *[]model.CommentPo) func() {
	return func() {
		defer wg.Done()
		select {
		case <-(*ctx).Done():
			return
		default:
			//TODO 远程调用order服务得到comments

		}

	}
}

func GetHouseImagesTask(ctx *context.Context, wg *sync.WaitGroup, houseId int, imgUrls *[]string) func() {
	return func() {
		defer wg.Done()
		select {
		case <-(*ctx).Done():
			return
		default:
			//获取house imgUrls
			urlsResp := GetHouseImgUrls(houseId)
			urls := make([]string, 0)
			json.Unmarshal(urlsResp.Data, &urls)
			for _, url := range urls {
				*imgUrls = append(*imgUrls, utils.ConcatImgUrl(conf.NginxUrl, url))
			}
			return
		}

	}
}

//SetESHouseTask  data:houseByte
func SetESHouseTask(data []byte) {
	esHousePo := &elasticsearch.ESHousePo{}
	utils.NewLog().Debug("house data:", string(data))
	//设置House字段
	house := &model.House{}
	json.Unmarshal(data, house)
	json.Unmarshal(data, esHousePo)

	esHousePo.HouseId = int64(house.ID)

	//user信息 web端做
	esHousePo.CreateTime = house.CreatedAt.Format(conf.MysqlTimeFormat)
	x, _ := json.Marshal(esHousePo)
	utils.NewLog().Debug("esHousePo:", string(x))
	ctx, _ := context.WithTimeout(context.Background(), conf.ESTaskTimeOut)
	esHousePos := []*elasticsearch.ESHousePo{esHousePo}
	err := elasticsearch.HouseES.BatchAdd(ctx, esHousePos)
	if err != nil {
		utils.NewLog().Info("elasticsearch.HouseES.BatchAdd error:", err)
	}
}
