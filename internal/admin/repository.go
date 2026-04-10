package admin

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

//   Admin
//   - GET /admin/reports — list reported content
//   - PUT /admin/reports/:id — resolve report (approve/dismiss)
//   - GET /admin/users — list users
//   - PUT /admin/users/:id/ban — ban user
//   - DELETE /admin/users/:id/ban — unban user
//   - GET /admin/ads — list ad slots
//   - POST /admin/ads — create ad
//   - PUT /admin/ads/:id — update ad
//   - DELETE /admin/ads/:id — delete ad

type AdminRepository interface {
	ListReports(context.Context) ([]Report, error)
	ResolveReport(context.Context, string) (*Report, error)
	ListUsers(context.Context) ([]User, error)
	BanUser(context.Context, string) (*User, error)
	UnbanUser(context.Context, string) (*User, error)
	ListAds(context.Context) ([]Ad, error)
	CreateAd(context.Context) (*Ad, error)
	UpdateAd(context.Context, string) (*Ad, error)
	DeleteAd(context.Context, string) error
}

type adminRepositoryImpl struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepositoryImpl{db: db}
}

func (r *adminRepositoryImpl) ListReports(ctx context.Context) ([]Report, error) {
	var reports []Report

	if err := r.db.WithContext(ctx).Find(&reports).Error; err != nil {
		return nil, fmt.Errorf("Cannot list reports")
	}

	return reports, nil
}

func (r *adminRepositoryImpl) ResolveReport(ctx context.Context, req ResolveReportRequest, id string) (*Report, error) {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Error; err != nil {
		return  nil, fmt.Errorf("Cannot find report")
	}

	
}