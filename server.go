package main

import (
	"chatroom/redis"
	"time"
)

func main() {
	redis.InitRedisPool("localhost:6379", 16, 0, 300*time.Second)
}