package model

import "ihome/service/user/utils"

func Register(phone, password string) error {
	mysqlPasswd := utils.Encryption(password)
	utils.NewLog().Info("mysqlPasswd:", mysqlPasswd)
	user := &User{Name: phone, Password_hash: mysqlPasswd}
	err := MysqlConn.Create(user).Error
	return err
}
