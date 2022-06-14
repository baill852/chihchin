package lib

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func NewRedis(config *Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.Db,
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Panic("Redis failed, ", err)
	}

	return client
}
