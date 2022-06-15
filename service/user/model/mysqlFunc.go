package model

import (
	"encoding/json"
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
		return utils.UserResponse(utils.RECODE_USERONERR, nil)
	}

	mysqlPasswd := utils.Encryption(password)
	utils.NewLog().Info("mysqlPasswd:", mysqlPasswd)
	user.Name = phone
	user.Password_hash = mysqlPasswd
	user.Mobile = phone
	err = MysqlConn.Create(user).Error
	if err != nil {
		utils.NewLog().Error("MysqlConn.Create error", err)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)
	}
	return utils.UserResponse(utils.RECODE_OK, nil)
}

func Login(phone, password string) kitex_gen.Response {
	user := User{}
	utils.NewLog().Info("Login MysqlConn:", MysqlConn.DB().Ping())
	//查询用户是否存在
	res := MysqlConn.Where("name = ?", phone).Or("mobile = ?", phone).First(&user)
	utils.NewLog().Info("MysqlConn.Where:", res.Error)
	if user.Name == "" {
		utils.NewLog().Info("not register:", user)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)

	}
	//判断密码是否正确
	if !utils.CheckPasswd(password, user.Password_hash) {
		utils.NewLog().Info("Password error:")
		return utils.UserResponse(utils.RECODE_PWDERR, nil)
	}
	//密码正确才登录成功,并保存登录session
	data, err := json.Marshal(user)
	if err != nil {
		utils.NewLog().Info("json.Marshal error:")
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)
	}

	return utils.UserResponse(utils.RECODE_OK, data)
}

//GetUserInfo 获取用户信息 sessionId->phone
func GetUserInfo(sessionId string) kitex_gen.Response {
	user := User{}
	utils.NewLog().Info("GetUserInfo MysqlConn:", MysqlConn.DB().Ping())
	//查询用户信息
	phone, _ := utils.AesEcpt.AesBase64Decrypt(sessionId)
	res := MysqlConn.Where("mobile = ?", phone).First(&user)
	utils.NewLog().Info("MysqlConn.Where:", res.Error)
	if user.Name == "" {
		utils.NewLog().Info("not register:", user)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)

	}
	data, err := json.Marshal(&user)
	if err != nil {
		utils.NewLog().Error("json.Marshal error", err)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)
	}
	return utils.UserResponse(utils.RECODE_OK, data)
}

func UpdateUserInfo(sessionId string, m map[string]string) kitex_gen.Response {
	utils.NewLog().Info("UpdateUserInfo MysqlConn:", MysqlConn.DB().Ping())
	//更新用户信息
	user := &User{}
	phone, _ := utils.AesEcpt.AesBase64Decrypt(sessionId)
	utils.NewLog().Info("info-value:", m)
	err := MysqlConn.Debug().Model(user).Where("mobile = ?", phone).Update(m).Error
	utils.NewLog().Info("MysqlConn.Where:", err)
	if err != nil {
		utils.NewLog().Error("update mysql fail:", err)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)
	}
	return utils.UserResponse(utils.RECODE_OK, nil)
}
