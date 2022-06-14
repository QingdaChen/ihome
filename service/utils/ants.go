package utils

import (
	"github.com/panjf2000/ants/v2"
	"ihome/service/user/conf"
	"time"
)

//协程池
type GoAntsPool struct {
	Pool *ants.Pool
	Err  error
}

var AntsPool *GoAntsPool

func init() {
	AntsPool = &GoAntsPool{}
	pool, err := ants.NewPool(conf.PoolSize,
		ants.WithExpiryDuration(time.Second*conf.PoolExpiryDuration),
		ants.WithPreAlloc(false),
	)
	if err != nil {
		NewLog().Error(" ants.NewPool error:", err)
	}
	AntsPool.Pool = pool
	AntsPool.Err = err
}
