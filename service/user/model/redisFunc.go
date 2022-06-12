package model

import (
	"context"
	"encoding/json"
	"errors"
	redis2 "github.com/go-redis/redis/v8"
	"ihome/service/user/conf"
	"ihome/service/user/kitex_gen"
	"ihome/service/utils"
	"strconv"
	"time"
)

var Client *redis2.Client

var ctx = context.Background()

func init() {
	Client = redis2.NewClient(&redis2.Options{
		Addr:        conf.RedisIp + ":" + strconv.Itoa(conf.RedisPort),
		Password:    "", // no password set
		DB:          0,  // use default DB
		PoolSize:    50,
		PoolTimeout: 5000, //池连接超时时间
	})

}

/*
  短信函数
*/
func SendSMSCode(phone, imgCode, uuid string) kitex_gen.Response {
	utils.NewLog().Info(phone + ":" + imgCode + ":" + uuid)
	conn := Client.Conn(ctx)
	//判断60s内是否已经发过短信
	result := ""
	err := errors.New("")
	result, err = conn.Get(ctx, phone+"_smsCode").Result()
	if result != "" {
		utils.NewLog().Error("SMSCode exists error:", err)
		return utils.UserResponse(utils.RECODE_SMSERR, nil)
	}

	//Client.Conn()
	result, err = conn.Get(ctx, uuid).Result()
	defer conn.Close()
	if err != nil {
		utils.NewLog().Error("SendSMSCode error", err)
		return utils.UserResponse(utils.RECODE_SMSERR, nil)
	}
	if result != imgCode {
		utils.NewLog().Error("SendSMSCode not equal error", result, imgCode)
		return utils.UserResponse(utils.RECODE_SMSEQERR, nil)
	}

	return utils.UserResponse(utils.RECODE_OK, nil)
}

func SaveSMSCode(phone, code string) kitex_gen.Response {
	conn := Client.Conn(ctx)
	defer conn.Close()
	_, err := conn.SetEX(ctx, phone+"_smsCode", code, conf.PhoneCodeTimeOut*time.Minute).Result()
	utils.NewLog().Info("save err:", err)
	if err != nil {
		utils.NewLog().Error("SaveSMSCode error:", err)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)
	}
	return utils.UserResponse(utils.RECODE_OK, nil)

}

func CheckSMSCode(phone, smsCode string) kitex_gen.Response {
	conn := Client.Conn(ctx)
	defer conn.Close()
	result, err := conn.Get(ctx, phone+"_smsCode").Result()
	utils.NewLog().Info("check err:", err)
	if err != nil {
		utils.NewLog().Error("CheckSMSCode error:", err)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)
	}
	if result != smsCode {
		utils.NewLog().Error("CheckSMSCode not equal:", err)
		return utils.UserResponse(utils.RECODE_SMSEQERR, nil)
	}

	return utils.UserResponse(utils.RECODE_OK, nil)
}

/*
   保存session
*/
func SaveRedisSession(data []byte) kitex_gen.Response {
	conn := Client.Conn(ctx)
	defer conn.Close()
	user := User{}
	err := json.Unmarshal(data, &user)
	if err != nil {
		utils.NewLog().Error("json.Unmarshal error:", err)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)
	}
	utils.NewLog().Info("user.Id:", user.Mobile)
	//将用户Id加密
	//utils.NewLog().Info("encrypt ", encrypt)
	phoneHash, _ := utils.AesEcpt.AesBase64Encrypt(user.Mobile)

	//直接存入用户名...
	_, err = conn.SetEX(ctx, conf.SessionLoginIndex+"_"+phoneHash, []byte(user.Name),
		conf.SessionLoginTimeOut*time.Hour).Result()
	utils.NewLog().Info("save err:", err)
	if err != nil {
		utils.NewLog().Info("conn.SetEX error:", err)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)
	}

	return utils.UserResponse(utils.RECODE_OK, []byte(phoneHash))

}

//CheckRedisSession session检查
func CheckRedisSession(sessionId string) kitex_gen.Response {
	conn := Client.Conn(ctx)
	defer conn.Close()
	utils.NewLog().Info("sessionId:", sessionId)
	result, err := conn.Exists(ctx, conf.SessionLoginIndex+"_"+sessionId).Result()
	utils.NewLog().Info("check result:", result)
	if err != nil {
		utils.NewLog().Error("conn.Exists error:", err)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)
	}
	if result == 0 {
		utils.NewLog().Info("user not login:", result)
		return utils.UserResponse(utils.RECODE_SESSIONERR, nil)
	}

	return utils.UserResponse(utils.RECODE_OK, nil)
}

//GetSessionInfo 获取session信息
func GetSessionInfo(sessionId string) kitex_gen.Response {
	conn := Client.Conn(ctx)
	defer conn.Close()
	utils.NewLog().Info("GetSessionInfo:", sessionId)
	result, err := conn.Get(ctx, conf.SessionLoginIndex+"_"+sessionId).Result()
	utils.NewLog().Info("GetSessionInfo result:", result)
	if result == "" {
		utils.NewLog().Info("GetSessionInfo nil:", err)
		return utils.UserResponse(utils.RECODE_SESSIONERR, nil)
	}
	user := &User{}
	user.Name = result
	data, _ := json.Marshal(user)
	return utils.UserResponse(utils.RECODE_OK, data)
}

//updateSessionInfo 更新session信息
func UpdateSessionInfo(sessionId string, updateName string) kitex_gen.Response {
	conn := Client.Conn(ctx)
	defer conn.Close()
	//直接存入用户名...
	_, err := conn.SetEX(ctx, conf.SessionLoginIndex+"_"+sessionId, []byte(updateName),
		conf.SessionLoginTimeOut*time.Hour).Result()
	utils.NewLog().Info("UpdateSessionInfo:", err)
	if err != nil {
		utils.NewLog().Info("conn.SetEX error:", err)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)
	}

	return utils.UserResponse(utils.RECODE_OK, nil)
}

//DeleteKey 根据sessionId删除session信息
func DeleteKey(key string) kitex_gen.Response {
	conn := Client.Conn(ctx)
	defer conn.Close()
	utils.NewLog().Info("DeleteKeyById:", key)
	_, err := conn.Del(ctx, key).Result()
	//utils.NewLog().Info("GetSessionInfo result:", result)
	if err != nil {
		utils.NewLog().Error("conn.Get error:", err)
		return utils.UserResponse(utils.RECODE_SERVERERR, nil)
	}

	return utils.UserResponse(utils.RECODE_OK, nil)
}

//GetRedisUserInfo 获取用户信息
func GetRedisUserInfo(sessionId string) kitex_gen.Response {
	conn := Client.Conn(ctx)
	defer conn.Close()
	utils.NewLog().Info("GetRedisUserInfo:", sessionId)
	result, err := conn.Get(ctx, conf.UserRedisIndex+"_"+sessionId).Result()
	utils.NewLog().Info("GetRedisUserInfo result:", result)
	if result == "" {
		utils.NewLog().Info("GetRedisUserInfo nil:", err)
		return utils.UserResponse(utils.RECODE_SESSIONERR, nil)
	}
	return utils.UserResponse(utils.RECODE_OK, []byte(result))
}

//SaveRedisUserInfo 保存用户信息
func SaveRedisUserInfo(sessionId string, data []byte) kitex_gen.Response {
	conn := Client.Conn(ctx)
	defer conn.Close()
	utils.NewLog().Info("SaveRedisUserInfo:", sessionId)
	_, err := conn.Set(ctx, conf.UserRedisIndex+"_"+sessionId, data, conf.UserInfoTimeOut*time.Hour).Result()
	utils.NewLog().Info("SaveRedisUserInfo result:", err)
	if err != nil {
		utils.NewLog().Info("SaveRedisUserInfo fail:", err)
		return utils.UserResponse(utils.RECODE_SESSIONERR, nil)
	}
	return utils.UserResponse(utils.RECODE_OK, nil)
}
