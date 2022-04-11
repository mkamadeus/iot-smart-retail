package entry

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func (s *Service) CheckIn(cardId *uuid.UUID) error {
	ctx := context.Background()
	status := s.Cache.Set(ctx, cardId.String(), time.Now().String(), 0)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}
