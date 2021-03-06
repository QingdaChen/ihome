package model

import (
	"github.com/Spongecaptain/redisLock"
	"github.com/go-redis/redis"
	"ihome/service/user/conf"
	"ihome/service/utils"
	"strconv"
	"time"
)

var redisClient *redis.Client
var Lock *redisLock.RedisLock

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:        conf.RedisIp + ":" + strconv.Itoa(conf.RedisPort),
		Password:    "", // no password set
		DB:          0,  // use default DB
		PoolSize:    50,
		PoolTimeout: 5000, //池连接超时时间
	})

}

func InitLock(key string) {
	Lock = redisLock.NewRedisLock(redisClient, key, conf.RedisLockValue)
}

//尝试释放锁
func TryRelease() {
	err := Lock.Unlock()
	if err != nil {
		utils.NewLog().Error("release redisLock fail:", err)
		//启动协程重试或发消息
		go releaseLock(Lock)
		return
	}
	utils.NewLog().Info("release redis lock success")
}

func releaseLock(lock *redisLock.RedisLock) {
	var i time.Duration = 1
	for {
		//每隔 1s 2s 4s重试
		err := Lock.Unlock()
		if err == nil || i > conf.RedisLockTimeOut {
			//发送成功或者大于32s返回
			break
		}
		time.Sleep(time.Second * i)
		i = i * 2
	}
}
