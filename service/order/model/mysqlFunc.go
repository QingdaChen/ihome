package model

import (
	"ihome/service/utils"
)

//GetMysqlOrder 获取订单
func GetMysqlOrder(houseId int) (*OrderHouse, error) {
	utils.NewLog().Printf("MysqlConn:%v", MysqlConn)
	order := &OrderHouse{}
	err := MysqlConn.Debug().Where("house_id = ?", houseId).
		First(order).Error
	if err != nil {
		utils.NewLog().Debug("GetMysqlAreas error:", err)
		return order, err
	}
	return order, nil

}

//CreateMysqlOrder 创建订单 Days ,Amount,Status
func CreateMysqlOrder(order *OrderHouse) error {
	utils.NewLog().Printf("MysqlConn:%v", MysqlConn)
	err := MysqlConn.Debug().Create(order).Error
	return err

}
