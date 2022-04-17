package entry

import (
	"context"
	"time"
)

func (s *Service) CheckIn(cardId string) error {
	ctx := context.Background()
	status := s.Cache.Set(ctx, cardId, time.Now().String(), 0)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}
