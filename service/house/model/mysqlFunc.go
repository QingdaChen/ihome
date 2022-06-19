package model

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"ihome/service/house/kitex_gen"
	"ihome/service/utils"
	"ihome/web/model"
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

func GetMysqlHouses(userId int) kitex_gen.Response {
	houses := make([]House, 0)
	err := MysqlConn.Debug().Where("user_id", userId).Find(&houses).Error
	if err != nil {
		utils.NewLog().Info("Find error:", err)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	data, _ := json.Marshal(&houses)
	return utils.HouseResponse(utils.RECODE_OK, data)

}

//GetMysqlHouseImg 获取house img
func GetMysqlHouseImg(houseId int) []string {
	type Url struct {
		Url string
	}
	houseImages := make([]Url, 0)
	res := make([]string, 0)
	err := MysqlConn.Debug().Model(&HouseImage{}).Order("id asc").
		Where("house_id = ?", houseId).Find(&houseImages).Error
	if err != nil {
		utils.NewLog().Info("Find error:", err)
		return res
	}
	for _, item := range houseImages {
		res = append(res, item.Url)
	}
	utils.NewLog().Debug("img urls res:", res)
	return res
}

//SaveMysqlHouseIdImg 保存houseID img
func SaveMysqlHouseIdImg(houseID int, imgUrl string) {
	err := MysqlConn.Debug().Model(&HouseImage{}).
		Create(&HouseImage{HouseId: uint(houseID), Url: imgUrl}).Error
	if err != nil {
		utils.NewLog().Info("SaveMysqlHouseIdImg error:", err)
	}

}

//GetMysqlHouseInfo 获取house信息
func GetMysqlHouseInfo(houseId int) (House, error) {
	house := House{}
	err := MysqlConn.Debug().Where("id = ?", houseId).First(&house).Error
	if err != nil {
		return house, err
	}
	return house, nil
}

func GetMysqlHouseFacilityIds(houseId int) []int {
	houseFacility := make([]HouseFacilities, 0)
	MysqlConn.Debug().Where("house_id=?", houseId).Find(&houseFacility)
	fid := make([]int, 0)
	for _, item := range houseFacility {
		fid = append(fid, item.FacilityId)
	}
	utils.NewLog().Debug("fid:", fid)
	return fid

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
	//id := strconv.Itoa(int(house.ID))
	data, _ := json.Marshal(house)
	return utils.HouseResponse(utils.RECODE_OK, data)
}

//GetMysqlHouseIds 获取用户的所有houseId
func GetMysqlHouseIds(userId int) *[]model.House {
	utils.NewLog().Debug("GetMysqlHouseIds...")
	houses := make([]model.House, 0)
	MysqlConn.Debug().Model(&House{}).Where("user_id=?", userId).Find(&houses)
	return &houses
}

func setHouse(house *House, m map[string]interface{}) {
	mapstructure.WeakDecode(m, house)
	house.UserId = uint(utils.PareInt(m["user_id"].(string)))
	house.AreaId = uint(utils.PareInt(m["area_id"].(string)))
	utils.NewLog().Debug("house:", house)
}
