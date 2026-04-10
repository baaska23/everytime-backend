package ads

import "context"

type Service struct {
	adRepo AdRepository
}

func NewService(adRepo AdRepository) *Service {
	return &Service{adRepo: adRepo}
}

func (s *Service) ListActiveBanners(ctx context.Context) ([]Ad, error) {
	return s.adRepo.ListActiveBanners(ctx)
}