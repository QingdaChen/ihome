package main

import (
	"context"
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
	"github.com/gomodule/redigo/redis"
	"ihome/service/captcha/conf"
	"ihome/web/utils"
	"strconv"
)

var ctx = context.Background()

func TestRedis() {
	conn, err := redis.Dial("tcp", "192.168.31.219:6379")
	if err != nil {
		fmt.Println("redis Dial err:", err)
		return
	}
	defer conn.Close()

	// 2. 操作数据库
	_, err = conn.Do("set", "itcast", "itheima")
}

func TestRedis2() {
	conn := redis2.NewClient(&redis2.Options{
		Addr:     conf.RedisIp + ":" + strconv.Itoa(conf.RedisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	utils.NewLog().Info("conn..", conn)
	//conn.SetEX(ctx, uuid, code, conf.RedisTimeOut)

}
