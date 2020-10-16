package redis

import (
	"WonderCV-MDR/conf"
	"fmt"
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func InitRedis() {
	fmt.Println("redisUrl :" + conf.REDIS_URL[8:])
	Redis = redis.NewClient(&redis.Options{
		Addr:       conf.REDIS_URL[8:],
		Password:   "", // no password set
		DB:         0,  // use default DB
		MaxRetries: 10,
		PoolSize:   100,
	})
	pong, err := Redis.Ping().Result()
	if err != nil {
		fmt.Printf("Redis : %v\n", err)
	} else {
		fmt.Printf("Redis : %v\n", pong)
	}

}
