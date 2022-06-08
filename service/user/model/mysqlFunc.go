package model

import (
	"ihome/service/user/kitex_gen"
	"ihome/service/utils"
)

func Register(phone, password string) kitex_gen.Response {
	//查询数据库判断用户是否已经注册
	user := &User{}
	utils.NewLog().Info("MysqlConn:", MysqlConn.DB().Ping())
	err := MysqlConn.Where("name = ?", phone).First(&user).Error
	utils.NewLog().Info("MysqlConn.Where:", err)

	//判定是否重复注册
	if "" != user.Name {
		utils.NewLog().Error("user registered:", err)
		return utils.UserResponse(utils.RECODE_USERONERR)
	}

	mysqlPasswd := utils.Encryption(password)
	utils.NewLog().Info("mysqlPasswd:", mysqlPasswd)
	user.Name = phone
	user.Password_hash = mysqlPasswd
	user.Mobile = phone
	err = MysqlConn.Create(user).Error
	if err != nil {
		utils.NewLog().Error("MysqlConn.Create error", err)
		return utils.UserResponse(utils.RECODE_SERVERERR)
	}
	return utils.UserResponse(utils.RECODE_OK)
}
