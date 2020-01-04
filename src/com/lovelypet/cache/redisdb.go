package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var client *redis.Client

func init() {
	fmt.Println("redis client init() called...")
	client = redis.NewClient(&redis.Options{
		Addr:"127.0.0.1:6379",
		Password:"",
		DB:0,
		PoolSize:10,
		MaxRetries:3,
		IdleTimeout:10*time.Second,
	})
	pong,err := client.Ping().Result()

	if err == redis.Nil {
		fmt.Println("redis 异常")
	}else if err != nil {
		fmt.Println("redis err:",err)
	}else {
		fmt.Println("redis init success:",pong)
	}
}

//获取自增唯一id
func IncrUniqueId(key string) int  {
	id,err := client.Incr(key).Result()
	if err != nil {
		fmt.Println("redis incr error:",err)
	}
	return int(id)
}



