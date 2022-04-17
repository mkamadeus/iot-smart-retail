package entry

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func (s *Service) Check(cardId *uuid.UUID) bool {
	ctx := context.Background()
	_, err := s.Cache.Get(ctx, cardId.String()).Result()
	if err == redis.Nil {
		return false
	}
	return true
}
