package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var redisPool *redis.Pool

func InitRedisPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	redisPool = &redis.Pool{
		MaxIdle: maxIdle,
		MaxActive: maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
