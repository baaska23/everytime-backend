package admin

import (
	"context"
	"fmt"
)

type ReportService interface {
	Create(ctx context.Context, req CreateReportRequest) (*Report, error)
	// GetByID(ctx context.Context, id uint) (*Report, error)
	// List(ctx context.Context, filter ReportFilter) (types.PagedResult[Report], error)
	// Resolve(ctx context.Context, id uint, req ResolveReportRequest) (*Report, error)
	// Delete(ctx context.Context, id uint) error
}

type reportServiceImpl struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) ReportService {
	return &reportServiceImpl{repo: repo}
}

func (s *reportServiceImpl) Create(ctx context.Context, req CreateReportRequest) (*Report, error) {
	if req.UserID == req.PostID {
		return nil, fmt.Errorf("cannot report your own post")
	}

	if req.Reason == "" {
		return nil, fmt.Errorf("reason must be provided")
	}

	existing, _ := s.repo.List(ctx, ReportFilter{
		UserID: req.UserID,
		PostID: req.PostID,
	})

	if existing.Total > 0 {
		return nil, fmt.Errorf("already reported")
	}
	
	return s.repo.Create(ctx, req)
}