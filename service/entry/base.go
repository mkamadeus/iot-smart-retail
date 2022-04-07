package entry

import (
	"github.com/go-redis/redis/v8"
)

type Service struct {
	Cache *redis.Client
}

func New(cache *redis.Client) *Service {
	return &Service{
		Cache: cache,
	}
}
