package models

import (
	"context"
	"errors"
	redis2 "github.com/go-redis/redis/v8"
	"ihome/service/captcha/conf"
	"ihome/web/utils"
	"strconv"
	"time"
)

var ctx = context.Background()

func InitRedis() *redis2.Client {
	client := redis2.NewClient(&redis2.Options{
		Addr:        conf.RedisIp + ":" + strconv.Itoa(conf.RedisPort),
		Password:    "", // no password set
		DB:          0,  // use default DB
		PoolSize:    50,
		PoolTimeout: 5000, //池连接超时时间
	})
	defer client.Close()
	return client
}
func SendSMSCd(phone, imgCode, uuid string) error {

	client := InitRedis()
	utils.NewLog().Info("phone...", phone)
	result, err := client.Get(ctx, uuid).Result()
	defer client.Close()
	if err != nil {
		return errors.New("SendSMSCd error")
	}
	if result != imgCode {
		return errors.New("SendSMSCd not equal")
	}

	return nil
}

func SaveSMSCd(phone, code string) error {
	client := InitRedis()
	_, err := client.SetEX(ctx, phone+"_"+code, code, conf.PhoneCodeTimeOut*time.Minute).Result()
	if err != nil {
		return errors.New("SaveSMSCd error")
	}
	return nil

}
