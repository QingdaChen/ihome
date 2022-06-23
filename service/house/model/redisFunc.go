package model

import (
	"context"
	redis2 "github.com/go-redis/redis/v8"
	"ihome/service/house/conf"
	"ihome/service/house/kitex_gen"
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

//GetRedis 获取redis中的信息
func GetRedis(key string) kitex_gen.Response {
	utils.NewLog().Info("GetRedis start...")
	conn := Client.Conn(ctx)
	result, _ := conn.Get(ctx, key).Result()
	//utils.NewLog().Info("GetRedis result...", result)
	response := utils.HouseResponse(utils.RECODE_OK, []byte(result))

	return response

}

//SaveRedis 保存信息到redis中
func SaveRedis(key string, data []byte, expire time.Duration) kitex_gen.Response {
	utils.NewLog().Info("SaveRedis start...")
	conn := Client.Conn(ctx)
	//redis保存areas
	_, err := conn.Set(ctx, key, data, expire).Result()
	if err != nil {
		utils.NewLog().Error("conn.Set error...", err)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	utils.NewLog().Debug("SaveRedis end...")
	return utils.HouseResponse(utils.RECODE_OK, nil)
}

//DeleteKey 根据key 删除redis缓存
func DeleteKey(key string) kitex_gen.Response {
	conn := Client.Conn(ctx)
	defer conn.Close()
	utils.NewLog().Info("DeleteKeyById:", key)
	_, err := conn.Del(ctx, key).Result()
	//utils.NewLog().Info("GetSessionInfo result:", result)
	if err != nil {
		utils.NewLog().Error("conn.Get error:", err)
		return utils.HouseResponse(utils.RECODE_SERVERERR, nil)
	}
	//conn.MGet()
	return utils.HouseResponse(utils.RECODE_OK, nil)
}
