package handler

import (
	"context"
	"encoding/json"
	"errors"
	"ihome/remote"
	conf_ihome "ihome/service/conf"
	model2 "ihome/service/model"
	"ihome/service/order/conf"
	"ihome/service/order/kitex_gen"
	"ihome/service/order/model"
	user_kitex_gen "ihome/service/user/kitex_gen"
	"ihome/service/utils"
)

var ctx, _ = context.WithTimeout(context.Background(), conf_ihome.RPCTimeOut)

//PostOrder 创建订单
//TODO 分布式锁
func PostOrder(req *kitex_gen.PostOrderReg) *kitex_gen.PostOrderResp {

	//获取User信息
	user, err := GetUserInfo(req.SessionId)
	resp := &kitex_gen.PostOrderResp{}
	if err != nil {
		//出错返回
		utils.NewLog().Debug("GetSession error:", err)
		resp.Errno = utils.RECODE_SERVERERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_SERVERERR)
		return resp
	}
	days := 0
	days, err = utils.GetDays(req.StartDate, req.EndDate)
	if err != nil {
		resp.Errno = utils.RECODE_SERVERERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_SERVERERR)
		return resp
	}
	//获取房源信息
	house, err2 := GetHouseInfo(int(req.HouseId))
	if err2 != nil {
		utils.NewLog().Debug("GetHouseInfo error:", err2)
		resp.Errno = utils.RECODE_SERVERERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_SERVERERR)
		return resp
	}
	if user.ID == int(house.UserId) {
		//房东不能预订自己发布房源
		resp.Errno = utils.RECODE_SERVERERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_SERVERERR)
		return resp
	}
	//判断日期是否冲突
	if !(house.Min_days <= days && house.Max_days >= days) {
		//住房日期不合适直接返回
		resp.Errno = utils.RECODE_SERVERERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_SERVERERR)
		return resp
	}
	//TODO 分布式锁
	model2.InitLock(utils.ConcatRedisKey(conf.OrderRedisLock, req.HouseId))
	model2.Lock.Lock()
	defer model2.Lock.Unlock()
	//确定房子当前没被预订
	orderMysql, _ := GetOrder(int(req.HouseId))
	if orderMysql.ID != 0 {
		//被预订了直接返回
		resp.Errno = utils.RECODE_ACCEPTED
		resp.Errmsg = utils.RecodeText(utils.RECODE_ACCEPTED)
		return resp
	}
	order := &model.OrderHouse{}
	//设置order
	SetOrder(order, req, days, user, house)
	//更新数据库
	err2 = model.CreateMysqlOrder(order)
	if err2 != nil {
		resp.Errno = utils.RECODE_SERVERERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_SERVERERR)
		return resp
	}
	//更新order redis
	data, _ := json.Marshal(order)
	model.SaveRedis(utils.ConcatRedisKey(conf.OrderHouseRedisKey, req.HouseId),
		data, conf.OrderRedisTimeOut)

	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	postResp := &kitex_gen.PostResp{OrderId: int64(order.ID)}
	resp.Data = postResp
	return resp
}

//GetOrder 获取订单
func GetOrder(houseId int) (*model.OrderHouse, error) {
	//先查redis
	orderRedis, _ := model.GetRedisOrder(utils.ConcatRedisKey(conf.OrderHouseRedisKey, houseId))
	if orderRedis.ID != 0 {
		//查到了返回
		return orderRedis, nil
	}
	//查不到查数据库
	orderMysql, _ := model.GetMysqlOrder(houseId)
	if orderMysql.ID != 0 {
		//已经被预订直接返回
		return orderMysql, errors.New("已经被预订")
	}

	return orderMysql, nil
}

//GetUserInfo 获取用户信息
func GetUserInfo(sessionId string) (*model2.UserPo, error) {
	user := &model2.UserPo{}
	req := user_kitex_gen.GetUserRequest{SessionId: sessionId}
	res, err := remote.RPC(ctx, conf.UserServiceIndex, req)
	if err != nil {
		utils.NewLog().Info("remote.RPC error:", err)
		return user, err
	}
	response := res.(*user_kitex_gen.Response)
	json.Unmarshal(response.Data, user)
	return user, nil
}
