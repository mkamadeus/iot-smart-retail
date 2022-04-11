package entry

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) CheckOut(cardId *uuid.UUID) error {
	ctx := context.Background()
	status := s.Cache.Del(ctx, cardId.String())
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}
