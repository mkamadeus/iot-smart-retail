package config

import "os"

type CacheConfig struct {
	Host     string
	Port     string
	Password string
}

func NewCacheConfig() *CacheConfig {
	return &CacheConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
}
