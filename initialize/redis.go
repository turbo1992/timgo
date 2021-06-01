package initialize

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func Redis()  {
	NewClient()
}

func NewClient() *redis.Client {
	if RedisClient != nil {
		return RedisClient
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		//Password: "123456", // no password set
		DB:   0, // use default DB
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println("redis connection failed!")
	}
	fmt.Println("redis connection success!")
	return RedisClient
}