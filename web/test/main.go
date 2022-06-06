package main

import (
	"ihome/service/user/utils"
)

func testEncryption() {
	mysqlPasswd := utils.Encryption("123456")
	utils.NewLog().Info("mysqlPasswd:", mysqlPasswd)
	utils.NewLog().Info("", utils.CheckPasswd("123456", mysqlPasswd))

}
func main() {

	//TestRedis2()
	//test()
	//TestMysql()
	testEncryption()

}
