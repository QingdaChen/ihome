package conf

import "time"

// user服务
const (
	ServerIp             string = "192.168.31.219"
	ServerPort           int    = 9002
	ServerMaxConnections int    = 1000
	ServerMaxQPS         int    = 100
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
