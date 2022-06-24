package conf

import "time"

//fastDfs
const (
	FastDfsCfgPath = "./conf/fastDfs.conf"
)

//RPC
const (
	RPCTimeOut = 4 * time.Second
)

//captcha
const (
	CaptchaServerIp     = "192.168.31.219"
	CaptchaServerPort   = 9000
	CaptchaServiceIndex = "captchaService"
)

//User
const (
	UserServiceIndex = "userService"
)

//House
const (
	HouseServiceIndex string = "houseService"
)

//order
const (
	NotOrdered = "WAIT_ACCEPT"
	Ordered    = "ACCEPTED"
)
