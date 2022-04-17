package entry

import (
	"context"
)

func (s *Service) CheckOut(cardId string) error {
	ctx := context.Background()
	status := s.Cache.Del(ctx, cardId)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}
