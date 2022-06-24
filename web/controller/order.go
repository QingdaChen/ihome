package controller

import (
	"github.com/gin-gonic/gin"
	"ihome/conf"
	"ihome/remote"
	order_kitex_gen "ihome/service/order/kitex_gen"
	"ihome/web/model"
	"ihome/web/utils"
	"net/http"
)

//PostOrders 创建订单
func PostOrders(ctx *gin.Context) {
	utils.NewLog().Info("PostOrders start...")
	orderReg := &model.PostOrderReq{}
	ctx.ShouldBind(orderReg)
	utils.NewLog().Debugf("houseId:%s start_date:%s end_date:%s", orderReg.HouseId,
		orderReg.StartDate, orderReg.EndDate)
	isOk := utils.CheckDate(ctx, orderReg.StartDate, orderReg.EndDate)
	if !isOk {
		return
	}
	//获得cookie
	sessionId, err := ctx.Cookie(conf.LoginCookieName)
	if err != nil || sessionId == "" {
		//sessionId 不存在或者过期直接返回
		utils.NewLog().Info("ctx.Cookie error:", err)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SESSIONERR, nil))
		return
	}
	//调用order服务
	req := order_kitex_gen.PostOrderReg{HouseId: int64(utils.PareInt(orderReg.HouseId)),
		SessionId: sessionId, StartDate: orderReg.StartDate, EndDate: orderReg.EndDate}
	res, err2 := remote.RPC(ctx, conf.OrderServerIndex, req)
	if err2 != nil {
		utils.NewLog().Info("remote.RPC error:", err2)
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SERVERERR, nil))
		return
	}
	//更新成功
	response := res.(*order_kitex_gen.PostOrderResp)
	utils.NewLog().Info("res:", response)

	ctx.JSON(http.StatusOK, response)
}
