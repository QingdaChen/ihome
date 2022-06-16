package main

import (
	"context"
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
	"github.com/tedcy/fdfs_client"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ihome/web/conf"
	"ihome/web/utils"
	"strconv"
	"time"
)

//func testEncryption() {
//	mysqlPasswd := utils.Encryption("123456")
//	utils.NewLog().Info("mysqlPasswd:", mysqlPasswd)
//	utils.NewLog().Info("", utils.CheckPasswd("123456", mysqlPasswd))
//
//}
func testRedis2() {
	conn := redis2.NewClient(&redis2.Options{
		Addr:     conf.SessionRedisIP + ":" + strconv.Itoa(conf.SessionRedisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	utils.NewLog().Info("conn..", conn)
	result, err := conn.Exists(context.Background(), "2").Result()
	utils.NewLog().Info("result..", result)
	utils.NewLog().Info("err:", err)
	//conn.SetEX(ctx, uuid, code, conf.RedisTimeOut)

}
func testCache() {
	//cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	//
	//cache.Set("my-unique-key", []byte("value"))
	//
	//entry, _ := cache.Get("my-unique-key")
	utils.InitCache()
	utils.SMSCache.Set("1", []byte("123"))
	time.Sleep(time.Second * 41)
	entry, _ := utils.SMSCache.Get("1")
	utils.NewLog().Info(string(entry))
}

//测试gin session
func testSession() {
	//r := gin.Default()
	//store, _ := redis.NewStore(10, "tcp",
	//	conf.SessionRedisIP+":"+strconv.Itoa(conf.SessionRedisPort), "", []byte("secret"))
	//r.Use(sessions.Sessions("mysession", store))

}

type HouseFacilities struct {
	HouseId    int `json:"house_id"`
	FacilityId int `json:"facility_id"`
}

func main() {

	//TestRedis2()
	//test()
	//TestMysql()
	//testEncryption()
	//testCache()
	//testRedis2()
	//testFastDfs()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/search_house?charset=utf8&parseTime=True&loc=Local",
		conf.MysqlUser, conf.MysqlPasswd, conf.MysqlIp, conf.MysqlPort)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.CreateInBatches(&[]HouseFacilities{{1, 2}, {1, 3}}, 2)

}

func testFastDfs() {
	client, err := fdfs_client.NewClientWithConfig("./fastDfs.conf")
	if err != nil {
		utils.NewLog().Error("fdfs_client init error", err)
	}
	filename, err2 := client.UploadByFilename("./test.jpg")
	utils.NewLog().Info("", filename, err2)
}
