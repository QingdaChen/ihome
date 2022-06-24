package handler

import (
	"encoding/json"
	"ihome/remote"
	"ihome/service/conf"
	house_kitex_gen "ihome/service/house/kitex_gen"
	"ihome/service/model"
	order_conf "ihome/service/order/conf"
	"ihome/service/order/kitex_gen"
	model2 "ihome/service/order/model"
	user_kitex_gen "ihome/service/user/kitex_gen"
	"ihome/service/utils"
)

//GetSession 根据sessionId获得session信息
func GetSession(sessionId string) (*model.SessionPo, error) {
	req := user_kitex_gen.SessionRequest{SessionId: sessionId}
	res, err := remote.RPC(remote.Ctx, conf.UserServiceIndex, req)
	session := &model.SessionPo{}
	if err != nil {
		utils.NewLog().Info("remote.RPC error:", err)
		return session, err
	}
	response := res.(*user_kitex_gen.Response)
	utils.NewLog().Debug("response:", response)
	json.Unmarshal(response.Data, session)
	return session, nil
}

func GetHouseInfo(houseId int) (*model.HouseRedisPo, error) {

	req := house_kitex_gen.GetHouseInfoReq{HouseId: int64(houseId)}
	res, err := remote.RPC(remote.Ctx, conf.HouseServiceIndex, req)
	houseInfo := &model.HouseRedisPo{}
	if err != nil {
		utils.NewLog().Info("remote.RPC error:", err)
		return houseInfo, err
	}
	response := res.(*house_kitex_gen.Response)
	utils.NewLog().Debug("response:", response)
	json.Unmarshal(response.Data, houseInfo)
	return houseInfo, nil

}

//更新ordr
func SetOrder(order *model2.OrderHouse, req *kitex_gen.PostOrderReg, days int,
	user *model.UserPo, house *model.HouseRedisPo) {
	order.UserId = uint(user.ID)
	order.HouseId = house.ID
	sd, _ := utils.TimeParse(req.StartDate)
	ed, _ := utils.TimeParse(req.EndDate)
	utils.NewLog().Debugf("sd:%v ed:%v", sd, ed)
	order.Begin_date = sd
	order.End_date = ed
	order.Days = days
	order.House_price = house.Price
	order.Amount = days * house.Price
	order.Status = order_conf.NotAccept
	//初始若有身份验证则个人征信良好
	if user.Id_card != "" {
		order.Credit = true
	} else {
		order.Credit = false
	}
}
