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

// UploadHouseImg implements the HouseServiceImpl interface.
func (s *HouseServiceImpl) UploadHouseImg(ctx context.Context, req *kitex_gen.UploadHouseImgReq) (resp *kitex_gen.Response, err error) {
	utils.NewLog().Debug("UploadHouseImg start....")
	uploadResp := handler.UploadHouseImg(int(req.HouseId), req.FileType, req.ImgBase64)
	if utils.RECODE_OK != uploadResp.Errno {
		utils.NewLog().Info("UploadHouseImg error", uploadResp)
	}
	return &uploadResp, nil
}

// GetHouseDetail implements the HouseServiceImpl interface.
func (s *HouseServiceImpl) GetHouseDetail(ctx context.Context, req *kitex_gen.GetHouseDetailReg) (resp *kitex_gen.HouseDetailResp, err error) {
	utils.NewLog().Debug("GetHouseDetail start...")
	detailResp := handler.GetHouseDetail(req.SessionId, int(req.HouseId))
	if utils.RECODE_OK != detailResp.Errno {
		utils.NewLog().Info("UploadHouseImg error", detailResp)
	}
	return detailResp, nil
}

// SearchHouse implements the HouseServiceImpl interface.
func (s *HouseServiceImpl) SearchHouse(ctx context.Context, req *kitex_gen.HouseSearchReq) (resp *kitex_gen.HouseSearchResp, err error) {
	utils.NewLog().Debug("SearchHouse start...")
	searchResp := handler.SearchHouse(req)
	if utils.RECODE_OK != searchResp.Errno {
		utils.NewLog().Info("SearchHouse error", searchResp)
	}
	return searchResp, nil
}

// HouseHomeIndex implements the HouseServiceImpl interface.
func (s *HouseServiceImpl) HouseHomeIndex(ctx context.Context, req *kitex_gen.HouseHomeIndexReg) (resp *kitex_gen.HouseSearchResp, err error) {
	utils.NewLog().Debug("HouseHomeIndex start...")
	indexResp := handler.GetHouseHomeIndex(req.SessionId)
	if utils.RECODE_OK != indexResp.Errno {
		utils.NewLog().Info("HouseHomeIndex error", indexResp)
	}
	return indexResp, nil
}
