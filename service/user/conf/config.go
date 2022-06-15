package conf

import "time"

// user服务
const (
	ServerIp             string = "192.168.31.219"
	ServerPort           int    = 9001
	UserRedisIndex       string = "userInfo"
	UserInfoTimeOut             = 30 * 24 //hour
	ServerMaxConnections int    = 1000
	ServerMaxQPS         int    = 100
	AvatarUrlIndex       string = "avatar_url"
)

//redis
const (
	RedisIp          string = "192.168.31.219"
	RedisPort        int    = 6379
	PhoneCodeTimeOut        = 1
	RedisLockKey     string = "reids-lock-key"
	RedisLockValue   string = "reids-lock-value"
	RedisLockTimeOut        = 4
)

//session
const (
	SessionRedisIP      string = "192.168.31.219"
	SessionRedisPort    int    = 6379
	SessionLoginIndex   string = "session_login"
	SessionLoginTimeOut        = 12 //hour
	SessionSecret       string = "asdggcgc"
)

//SMS
const (
	PhoneCode string = "123456"
)

//fastDfs
const (
	FastDfsCfgFilePath string = "conf/fastDfs.conf"
)

//nginx
const (
	NginxUrl string = "http://192.168.31.219:8888"
)

//协程池子
const (
	PoolSize           int           = 100
	PoolExpiryDuration time.Duration = 30
)
