package model

import (
	"context"
	redis2 "github.com/go-redis/redis/v8"
	"ihome/service/house/conf"
	"ihome/service/house/kitex_gen"
	"ihome/service/utils"
	"strconv"
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

//GetRedisAreas 获取redis中的areas
func GetRedisAreas() kitex_gen.Response {
	utils.NewLog().Info("GetRedisAreas start...")
	conn := Client.Conn(ctx)
	result := ""
	result, _ = conn.Get(ctx, conf.RedisAreasIndex).Result()
	utils.NewLog().Info("GetRedisAreas result...", result)
	response := utils.HouseResponse(utils.RECODE_OK, []byte(result))
	return response

}

//SaveRedisAreas 保存area到redis中
func SaveRedisAreas(areas []byte) kitex_gen.Response {
	utils.NewLog().Info("SaveRedisAreas start...")
	conn := Client.Conn(ctx)
	//redis保存areas
	_, err := conn.Set(ctx, conf.RedisAreasIndex, areas, conf.RedisAreasTimeOut).Result()
	if err != nil {
		utils.NewLog().Error("conn.Set error...", err)
		return utils.HouseResponse(utils.RECODE_SERVERERR, []byte(""))
	}
	return utils.HouseResponse(utils.RECODE_OK, []byte(""))

}
