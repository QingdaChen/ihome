package cache

import (
	"github.com/allegro/bigcache/v3"
	"ihome/service/utils"
	"time"
)

var FacilityCache *bigcache.BigCache

func init() {
	//初始化区域缓存
	config := bigcache.Config{
		Shards:             1024,
		LifeWindow:         65535 * time.Hour, //永远不过期
		CleanWindow:        0,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
		OnRemove:           nil,
		OnRemoveWithReason: nil,
	}
	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		utils.NewLog().Fatal("Cache init error", err)
	}
	FacilityCache = cache

}
