package app

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func InitializeCache(config Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Cache.Host, config.Cache.Port),
		Password: config.Cache.Password,
	})

	return rdb
}
