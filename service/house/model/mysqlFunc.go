package model

import (
	"encoding/json"
	"ihome/service/house/kitex_gen"
	"ihome/service/utils"
	"strconv"
)

//GetMysql obj 使用对象指针: obj := &[]Area{}
func GetMysql(obj interface{}) kitex_gen.Response {
	//area := &[]Area{}
	utils.NewLog().Printf("MysqlConn:%v", MysqlConn)
	err := MysqlConn.Find(obj).Error
	if err != nil {
		utils.NewLog().Error("GetMysqlAreas error:", err)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	//utils.NewLog().Info("obj:", obj)
	result, err2 := json.Marshal(obj)
	if err2 != nil {
		utils.NewLog().Error("json.Marshal error:", err2)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	//utils.NewLog().Debug("result:", result)
	return utils.HouseResponse(utils.RECODE_OK, result)

}

func SaveMysqlHouseFac(houseFacs []HouseFacilities) kitex_gen.Response {
	utils.NewLog().Info("HouseFacilities:", houseFacs)
	//sql := ""
	//i := 0
	//for i = 0; i < len(houseFacs)-1; i++ {
	//	sql = fmt.Sprintf("%s(%d,%d),", sql, houseFacs[i].HouseId, houseFacs[i].FacilityId)
	//}
	//sql = fmt.Sprintf("%s(%d,%d)", sql, houseFacs[i].HouseId, houseFacs[i].FacilityId)
	//
	//sql = fmt.Sprintf("INSERT INTO house_facilities VALUES %s;", sql)
	//utils.NewLog().Info("sql: ", sql)
	//result := MysqlConn.Debug().Exec(sql)
	result := MysqlConn.Debug().Create(&houseFacs)
	err := result.Error
	if err != nil {
		utils.NewLog().Info("Create error:", err)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	utils.NewLog().Info("SaveMysqlHouseFac success!!!")
	return utils.HouseResponse(utils.RECODE_OK, nil)
}

func SaveMysqlHouse(houseMap map[string]interface{}) kitex_gen.Response {
	house := &House{}
	mapData, _ := json.Marshal(&houseMap)
	json.Unmarshal(mapData, house)
	//houseMap["Facilities"] = ""
	utils.NewLog().Info("house:", house)
	result := MysqlConn.Debug().Create(house)
	err := result.Error
	if err != nil {
		utils.NewLog().Info("mysql Create:", err)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	id := strconv.Itoa(int(house.ID))
	return utils.HouseResponse(utils.RECODE_OK, []byte(id))
}
