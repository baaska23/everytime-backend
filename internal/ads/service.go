package ads

import (
	"context"
	"everytime-backend/internal/shared/types"
)

type Service struct {
	adRepo AdRepository
}

func NewService(adRepo AdRepository) *Service {
	return &Service{adRepo: adRepo}
}

func (s *Service) ListActiveBanners(ctx context.Context) ([]Ad, error) {
	return s.adRepo.ListActiveBanners(ctx)
}

func (s *Service) Interleave(ctx context.Context, items types.ListEntry, every init) ([]types.ListEntry, error)