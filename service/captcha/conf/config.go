package conf

// captcha服务
const (
	ServerIp             = "192.168.31.219"
	ServerPort           = 9000
	ServerMaxConnections = 1000
	ServerMaxQPS         = 100
)

//redis
const (
	RedisIp          = "192.168.31.219"
	RedisPort        = 6379
	RedisTimeOut     = 5 //second
	PhoneCodeTimeOut = 3
)
