package model

import(
	"github.com/gomodule/redigo/redis"
	"sync"
)

// [FILE FUNCTION]
// initialize

var Pool *redis.Pool

type model struct {
	pool *redis.Pool
}

var instance *model
var lock sync.Mutex

// get Model instance which is thread safety
// out: unique model instance
func GetModel() *model {
	// double check
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &model{
				pool: Pool,
			}
		}
	}
	return instance
}