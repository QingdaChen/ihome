package handler

import (
	"context"
	"encoding/json"
	"ihome/service/house/model"
	po "ihome/service/model"
	"ihome/service/utils"
	"ihome/web/conf"
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
