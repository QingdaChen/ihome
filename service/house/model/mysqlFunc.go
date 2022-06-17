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
	result := MysqlConn.Debug().Create(&houseFacs)
	err := result.Error
	if err != nil {
		utils.NewLog().Info("Create error:", err)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	utils.NewLog().Info("SaveMysqlHouseFac success!!!")
	return utils.HouseResponse(utils.RECODE_OK, nil)
}

func GetMysqlHouse(userId int) kitex_gen.Response {
	houses := make([]House, 0)
	err := MysqlConn.Debug().Where("user_id", userId).Find(&houses).Error
	if err != nil {
		utils.NewLog().Info("Find error:", err)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	data, _ := json.Marshal(&houses)
	return utils.HouseResponse(utils.RECODE_OK, data)

}
func SaveMysqlHouse(houseMap map[string]interface{}) kitex_gen.Response {
	utils.NewLog().Info("houseMap:", houseMap)
	house := &House{}
	setHouse(house, houseMap)
	//houseMap["Facilities"] = ""
	utils.NewLog().Info("house:", house)
	result := MysqlConn.Debug().Model(&House{}).Create(house)
	err := result.Error
	if err != nil {
		utils.NewLog().Info("mysql Create:", err)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	id := strconv.Itoa(int(house.ID))
	return utils.HouseResponse(utils.RECODE_OK, []byte(id))
}

func setHouse(house *House, m map[string]interface{}) {
	userId, _ := strconv.ParseInt(m["user_id"].(string), 10, 32)
	house.UserId = uint(userId)
	areaId, _ := strconv.ParseInt(m["area_id"].(string), 10, 32)
	house.AreaId = uint(areaId)
	house.Title = m["title"].(string)
	house.Address = m["address"].(string)
	roomCount, _ := strconv.ParseInt(m["room_count"].(string), 10, 32)
	house.Room_count = int(roomCount)
	acreage, _ := strconv.ParseInt(m["acreage"].(string), 10, 32)
	house.Acreage = int(acreage)
	price, _ := strconv.ParseInt(m["price"].(string), 10, 32)
	house.Price = int(price)
	house.Unit = m["unit"].(string)
	utils.NewLog().Info("capacity:", m["capacity"])
	capacity, _ := strconv.ParseInt(m["capacity"].(string), 10, 32)
	house.Capacity = int(capacity)
	house.Beds = m["beds"].(string)
	deposit, _ := strconv.ParseInt(m["deposit"].(string), 10, 32)
	house.Deposit = int(deposit)
	minDays, _ := strconv.ParseInt(m["min_days"].(string), 10, 32)
	house.Min_days = int(minDays)
	maxDays, _ := strconv.ParseInt(m["max_days"].(string), 10, 32)
	house.Max_days = int(maxDays)
	//orderCount, _ := strconv.ParseInt(m["order_count"].(string), 10, 32)
	//house.Order_count = int(orderCount)
	//house.Index_image_url = m["index_image_url"].(string)
	//mapstructure.Decode(m, house)
}
