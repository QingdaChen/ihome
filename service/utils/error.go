package utils

import (
	"ihome/conf"
	house_kitex_gen "ihome/service/house/kitex_gen"
	user_kitex_gen "ihome/service/user/kitex_gen"
)

const (
	RECODE_OK        = "0"
	RECODE_DBERR     = "4001"
	RECODE_NODATA    = "4002"
	RECODE_DATAEXIST = "4003"
	RECODE_DATAERR   = "4004"

	RECODE_SESSIONERR = "4101"
	RECODE_LOGINERR   = "4102"
	RECODE_PARAMERR   = "4103"
	RECODE_USERONERR  = "4104"
	RECODE_ROLEERR    = "4105"
	RECODE_PWDERR     = "4106"
	RECODE_USERERR    = "4107"
	RECODE_SMSERR     = "4108"
	RECODE_MOBILEERR  = "4109"
	RECODE_SMSEQERR   = "4110"

	RECODE_REQERR    = "4201"
	RECODE_IPERR     = "4202"
	RECODE_THIRDERR  = "4301"
	RECODE_IOERR     = "4302"
	RECODE_SERVERERR = "4500"
	RECODE_UNKNOWERR = "4501"
	RECODE_ACCEPTED  = "4502"
)

var recodeText = map[string]string{
	RECODE_OK:         "成功",
	RECODE_DBERR:      "数据库查询错误",
	RECODE_NODATA:     "无数据",
	RECODE_DATAEXIST:  "数据已存在",
	RECODE_DATAERR:    "数据错误",
	RECODE_SESSIONERR: "用户未登录",
	RECODE_LOGINERR:   "用户登录失败",
	RECODE_PARAMERR:   "参数错误",
	RECODE_USERERR:    "用户不存在或未激活",
	RECODE_USERONERR:  "用户已经注册",
	RECODE_ROLEERR:    "用户身份错误",
	RECODE_PWDERR:     "密码错误",
	RECODE_REQERR:     "非法请求或请求次数受限",
	RECODE_IPERR:      "IP受限",
	RECODE_THIRDERR:   "第三方系统错误",
	RECODE_IOERR:      "文件读写错误",
	RECODE_SERVERERR:  "内部错误",
	RECODE_UNKNOWERR:  "未知错误",
	RECODE_SMSERR:     "短信失败",
	RECODE_MOBILEERR:  "手机号错误",
	RECODE_SMSEQERR:   "短信与图像验证码不相等",
	RECODE_ACCEPTED:   "房子已经被预订",
}

func RecodeText(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}

func UserResponse(errCode string, data []byte) user_kitex_gen.Response {
	return user_kitex_gen.Response{Errno: errCode, Errmsg: RecodeText(errCode), Data: data}
}

func HouseResponse(errCode string, data []byte) house_kitex_gen.Response {
	return house_kitex_gen.Response{Errno: errCode, Errmsg: RecodeText(errCode), Data: data}
}

func HouseDetailResponse(errCode string, data *house_kitex_gen.HouseDetailData) house_kitex_gen.HouseDetailResp {
	return house_kitex_gen.HouseDetailResp{Errno: errCode, Errmsg: RecodeText(errCode), Data: data}
}

func Response(code string, data interface{}) map[string]interface{} {
	resp := make(map[string]interface{}, 3)
	resp[conf.ErrorNoIndex] = code
	resp[conf.ErrorMsgIndex] = RecodeText(code)
	resp[conf.DataIndex] = data
	return resp

}
