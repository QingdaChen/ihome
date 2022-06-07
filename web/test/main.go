package main

import (
	"github.com/allegro/bigcache/v3"
	"ihome/service/user/utils"
	"time"
)

func testEncryption() {
	mysqlPasswd := utils.Encryption("123456")
	utils.NewLog().Info("mysqlPasswd:", mysqlPasswd)
	utils.NewLog().Info("", utils.CheckPasswd("123456", mysqlPasswd))

}

func testCache() {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))

	cache.Set("my-unique-key", []byte("value"))

	entry, _ := cache.Get("my-unique-key")
	utils.NewLog().Info(string(entry))
}
func main() {

	//TestRedis2()
	//test()
	//TestMysql()
	//testEncryption()
	testCache()

}
