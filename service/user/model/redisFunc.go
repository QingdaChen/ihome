package model

import (
	"context"
	"errors"
	redis2 "github.com/go-redis/redis/v8"
	"ihome/service/user/conf"
	"ihome/service/user/kitex_gen"
	"ihome/service/user/utils"
	"strconv"
	"time"
)

var Client *redis2.Client

var ctx = context.Background()

func InitRedis() {
	Client = redis2.NewClient(&redis2.Options{
		Addr:        conf.RedisIp + ":" + strconv.Itoa(conf.RedisPort),
		Password:    "", // no password set
		DB:          0,  // use default DB
		PoolSize:    50,
		PoolTimeout: 5000, //池连接超时时间
	})

}

//func InitPool() {
//	Pool = &redis.Pool{
//		// Maximum number of connections allocated by the pool at a given time.
//		// When zero, there is no limit on the number of connections in the pool.
//		//最大活跃连接数，0代表无限
//		MaxActive: 888,
//		//最大闲置连接数
//		// Maximum number of idle connections in the pool.
//		MaxIdle: 20,
//		//闲置连接的超时时间
//		// Close connections after remaining idle for this duration. If the value
//		// is zero, then idle connections are not closed. Applications should set
//		// the timeout to a value less than the server's timeout.
//		IdleTimeout: time.Second * 100,
//		//定义拨号获得连接的函数
//		// Dial is an application supplied function for creating and configuring a
//		// connection.
//		//
//		// The connection returned from Dial must not be in a special state
//		// (subscribed to pubsub channel, transaction started, ...).
//		Dial: func() (redis.Conn, error) {
//			return redis.Dial("tcp", conf.RedisIp+":"+strconv.Itoa(conf.RedisPort))
//		},
//	}
//	//延迟关闭连接池
//	defer Pool.Close()
//
//}

//func SaveImgCode(uuid, code string, pool *redis.Pool) error {
//	//获取redis连接池连接
//	utils.NewLog().Info("pool..", pool)
//	conn := Pool.Get()
//	utils.NewLog().Info("conn..", conn)
//	defer conn.Close()
//	// 2. 操作数据库
//	_, err := conn.Do("setex", uuid, conf.RedisTimeOut, code)
//
//	return err
//
//}

func SendSMSCode(phone, imgCode, uuid string) kitex_gen.Response {
	utils.NewLog().Info(phone + ":" + imgCode + ":" + uuid)
	conn := Client.Conn(ctx)
	//判断60s内是否已经发过短信
	result := ""
	err := errors.New("")
	result, err = conn.Get(ctx, phone+"_smsCode").Result()
	if result != "" {
		utils.NewLog().Error("SMSCode exists error:", err)
		return utils.UserResponse(utils.RECODE_SMSERR)
	}

	//Client.Conn()
	result, err = conn.Get(ctx, uuid).Result()
	defer conn.Close()
	if err != nil {
		utils.NewLog().Error("SendSMSCode error", err)
		return utils.UserResponse(utils.RECODE_SMSERR)
	}
	if result != imgCode {
		utils.NewLog().Error("SendSMSCode not equal error", result, imgCode)
		return utils.UserResponse(utils.RECODE_SMSEQERR)
	}

	return utils.UserResponse(utils.RECODE_OK)
}

func SaveSMSCode(phone, code string) kitex_gen.Response {
	conn := Client.Conn(ctx)
	defer conn.Close()
	_, err := conn.SetEX(ctx, phone+"_smsCode", code, conf.PhoneCodeTimeOut*time.Minute).Result()
	utils.NewLog().Info("save err:", err)
	if err != nil {
		utils.NewLog().Error("SaveSMSCode error:", err)
		return utils.UserResponse(utils.RECODE_SERVERERR)
	}
	return utils.UserResponse(utils.RECODE_OK)

}

func CheckSMSCode(phone, smsCode string) kitex_gen.Response {
	conn := Client.Conn(ctx)
	defer conn.Close()
	result, err := conn.Get(ctx, phone+"_smsCode").Result()
	utils.NewLog().Info("check err:", err)
	if err != nil {
		utils.NewLog().Error("CheckSMSCode error:", err)
		return utils.UserResponse(utils.RECODE_SERVERERR)
	}
	if result != smsCode {
		utils.NewLog().Error("CheckSMSCode not equal:", err)
		return utils.UserResponse(utils.RECODE_SMSEQERR)
	}

	return utils.UserResponse(utils.RECODE_OK)
}
