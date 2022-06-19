package conf

//mysql
const (
	MysqlUser   string = "root"
	MysqlPasswd string = "220108"
	MysqlIp     string = "192.168.31.219"
	MysqlPort   int    = 3306
)

// web
const (
	WebPort           int = 8088
	UploadFileMaxSize     = 1024 * 1024 * 5 //5M
)

// captcha服务
const (
	CaptchaServerIp     string = "192.168.31.219"
	CaptchaServerPort   int    = 9000
	CaptchaServiceIndex string = "captchaService"
)

//user服务
const (
	UserServerIp     string = "192.168.31.219"
	UserServerPort   int    = 9001
	UserServiceIndex string = "userService"
)

//house 服务
const (
	HouseServerIp        string = "192.168.31.219"
	HouseServerPort      int    = 9002
	HouseAreasCacheIndex string = "areasData"
	HouseServiceIndex    string = "houseService"
)

//session
const (
	SessionRedisIP    string = "192.168.31.219"
	SessionRedisPort  int    = 6379
	SessionSize       int    = 10
	SessionNetwork    string = "tcp"
	SessionPasswd     string = ""
	SessionSecret     string = "2332456"
	SessionLoginIndex string = "session_login"
)

//cookie
const (
	LoginCookieName    string = "sessionId" //登录cookie name
	LoginCookieTimeOut int    = 3600 * 24 * 7
)

//web json返回
const (
	ErrorNoIndex  string = "errno"
	ErrorMsgIndex string = "errmsg"
	DataIndex     string = "data"
)

// LoginHtmlLocation file location
const (
	LoginHtmlLocation string = "http://192.168.31.219:8088/home/login.html"
)

//nginx
const (
	NginxUrl string = "http://192.168.31.219:8888"
)

var AccessId string = "LTAI5tCuiKcEcUEoJkoXZxQX"
var AccessSecret string = "aMSOHFuPS3vXhHyAJPnpoL32qFOaGw"
