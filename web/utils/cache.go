package utils

import (
	"github.com/allegro/bigcache/v3"
	"time"
)

var SMSCache *bigcache.BigCache
var AreasCache *bigcache.BigCache

func InitCache() {
	config := bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,
		// time after which entry can be evicted
		LifeWindow: 20 * time.Second,
		// Interval between removing expired en tries (clean up).
		// If set to <= 0 then no action is performed.
		// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
		CleanWindow: 20 * time.Second,
		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,
		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,
		// prints information about additional memory allocation
		Verbose: true,
		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 8192,
		// callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,
		// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A constant representing the reason will be passed through.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		// Ignored if OnRemove is specified.
		OnRemoveWithReason: nil,
	}
	cache, initErr := bigcache.NewBigCache(config)
	if initErr != nil {
		NewLog().Fatal("Cache init error", initErr)
	}
	SMSCache = cache
	//初始化区域缓存
	config = bigcache.Config{
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
	cache, initErr = bigcache.NewBigCache(config)
	if initErr != nil {
		NewLog().Fatal("Cache init error", initErr)
	}
	AreasCache = cache

}
