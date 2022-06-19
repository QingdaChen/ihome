package conf

import "time"

// house服务
const (
	ServerIp              string = "192.168.31.219"
	ServerPort            int    = 9002
	ServerMaxConnections  int    = 1000
	ServerMaxQPS          int    = 100
	HouseInfoRedisIndex   string = "houseInfo"
	HouseInfoRedisTimeOut        = 65535 * time.Hour //hour
	HouseImgRedisIndex    string = "houseImages"
	HouseImgRedisTimeOut         = 24 * 30 * time.Hour
	HouseTaskTimeOut             = 5 * time.Second
)

// user服务
const (
	UserServerIp             string = "192.168.31.219"
	UserServerPort           int    = 9001
	UserServiceIndex         string = "userService"
	UserHousesRedisIndex     string = "user_housesInfo"
	UserRedisIndex           string = "userInfo"
	UserHousesRedisTimeOut          = 2 * time.Hour //hour
	UserServerMaxConnections int    = 1000
	UserServerMaxQPS         int    = 100
	AvatarUrlIndex           string = "avatar_url"
)

const (
	FacilityIndex        string = "facilityData"
	FacilityRedisTimeOut        = 65535 * time.Hour
)

//redis
const (
	RedisIp           = "192.168.31.219"
	RedisPort         = 6379
	RedisAreasIndex   = "areaData"
	RedisAreasTimeOut = 65535 * time.Hour //2小时
)

//mysql
const (
	MysqlUser       string = "root"
	MysqlPasswd     string = "220108"
	MysqlIp         string = "192.168.31.219"
	MysqlPort       int    = 3306
	MysqlTimeFormat string = "2006-01-02 15:04:05"
)

//session
const (
	SessionRedisIP      string = "192.168.31.219"
	SessionRedisPort    int    = 6379
	SessionLoginIndex   string = "session_login"
	SessionLoginTimeOut        = 12 //hour
	SessionSecret       string = "asdggcgc"
)

//nginx
const (
	NginxUrl string = "http://192.168.31.219:8888"
)
