package conf

import "time"

// user服务
const (
	ServerIp   = "192.168.31.219"
	ServerPort = 9002
)

//redis
const (
	RedisIp           = "192.168.31.219"
	RedisPort         = 6379
	RedisAreasIndex   = "areaData"
	RedisAreasTimeOut = 2 * time.Hour //2小时
)

//mysql
const (
	MysqlUser   string = "root"
	MysqlPasswd string = "220108"
	MysqlIp     string = "192.168.31.219"
	MysqlPort   int    = 3306
)
