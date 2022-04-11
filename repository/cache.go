package repository

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/mkamadeus/iot-smart-retail/config"
)

func NewCache(c *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.Cache.Host, c.Cache.Port),
		Password: c.Cache.Password,
	})

	return rdb
}
