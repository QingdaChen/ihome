package conf

//web
const (
	WebPort int = 8088
)

// captcha服务
const (
	CaptchaServerIp   string = "192.168.31.219"
	CaptchaServerPort int    = 9000
)

//user
const (
	UserServerIp   string = "192.168.31.219"
	UserServerPort int    = 9001
)

//mysql
const (
	MysqlUser   string = "root"
	MysqlPasswd string = "220108"
	MysqlIp     string = "192.168.31.219"
	MysqlPort   int    = 3306
)

//house
const (
	HouseServerIp        string = "192.168.31.219"
	HouseServerPort      int    = 9002
	HouseAreasCacheIndex string = "areasData"
)

var AccessId string = "LTAI5tCuiKcEcUEoJkoXZxQX"
var AccessSecret string = "aMSOHFuPS3vXhHyAJPnpoL32qFOaGw"
