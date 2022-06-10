package conf

// user服务
const (
	ServerIp   string = "192.168.31.219"
	ServerPort int    = 9001
)

//redis
const (
	RedisIp          string = "192.168.31.219"
	RedisPort        int    = 6379
	PhoneCodeTimeOut        = 1
)

//session
const (
	SessionRedisIP      string = "192.168.31.219"
	SessionRedisPort    int    = 6379
	SessionLoginIndex   string = "session_login"
	SessionLoginTimeOut        = 12 //hour
)
