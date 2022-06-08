package model

import (
	"encoding/json"
	"ihome/service/house/kitex_gen"
	"ihome/service/utils"
)

func GetMysqlAreas() kitex_gen.Response {
	area := &[]Area{}
	utils.NewLog().Info("MysqlConn:", MysqlConn.DB().Ping())
	err := MysqlConn.Find(area).Error
	if err != nil {
		utils.NewLog().Error("GetMysqlAreas error:", err)
		return utils.HouseResponse(utils.RECODE_SERVERERR, []byte(""))
	}
	result, err2 := json.Marshal(area)
	if err2 != nil {
		utils.NewLog().Error("json.Marshal error:", err2)
		return utils.HouseResponse(utils.RECODE_SERVERERR, []byte(""))
	}
	return utils.HouseResponse(utils.RECODE_OK, result)

}
