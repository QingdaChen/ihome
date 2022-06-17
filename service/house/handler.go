package main

import (
	"context"
	"encoding/json"
	"ihome/service/house/handler"
	"ihome/service/house/kitex_gen"
	"ihome/service/utils"
)

// HouseServiceImpl implements the last service interface defined in the IDL.
type HouseServiceImpl struct{}

// GetArea implements the HouseServiceImpl interface.
func (s *HouseServiceImpl) GetArea(ctx context.Context, req *kitex_gen.AreaRequest) (resp *kitex_gen.Response, err error) {
	utils.NewLog().Debug("GetArea start....")
	areasResp := handler.GetAreas()
	if utils.RECODE_OK != areasResp.Errno {
		utils.NewLog().Info("areasResp error:", areasResp)
	}
	return &areasResp, nil
}

// PubHouse implements the HouseServiceImpl interface.
func (s *HouseServiceImpl) PubHouse(ctx context.Context, req *kitex_gen.PubHouseRequest) (resp *kitex_gen.Response, err error) {
	utils.NewLog().Debug("PubHouse start...")
	houseMap := make(map[string]interface{}, 10)
	err = json.Unmarshal(req.Params, &houseMap)
	if err != nil {
		utils.NewLog().Info("json.Unmarshal error:", err)
		response := utils.HouseResponse(utils.RECODE_SERVERERR, nil)
		return &response, nil
	}
	pubResp := handler.PubHouse(req.SessionId, houseMap)
	if utils.RECODE_OK != pubResp.Errno {
		utils.NewLog().Info("PubHouse error:", pubResp)
	}
	return &pubResp, nil

}

// GetUserHouse implements the HouseServiceImpl interface.
func (s *HouseServiceImpl) GetUserHouse(ctx context.Context, req *kitex_gen.GetUserHouseRequest) (resp *kitex_gen.Response, err error) {
	utils.NewLog().Debug("GetHouse start....")
	userHouseResp := handler.GetUserHouse(req.SessionId)
	if utils.RECODE_OK != userHouseResp.Errno {
		utils.NewLog().Info("GetUserHouse error:", userHouseResp)
	}
	return &userHouseResp, nil
}
