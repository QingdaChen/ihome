package main

import (
	"context"
	"ihome/service/order/handler"
	"ihome/service/order/kitex_gen"
	"ihome/service/utils"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PostOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PostOrder(ctx context.Context, req *kitex_gen.PostOrderReg) (resp *kitex_gen.PostOrderResp, err error) {
	utils.NewLog().Debug("PostOrder start...")
	orderResp := handler.PostOrder(req)
	if utils.RECODE_OK != orderResp.Errno {
		utils.NewLog().Info("HouseHomeIndex error", orderResp)
	}
	return orderResp, nil
}
