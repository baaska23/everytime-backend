package admin

import (
	"context"
	"everytime-backend/internal/shared/types"
	"fmt"

	"gorm.io/gorm"
)

type ReportRepository interface {
	Create(ctx context.Context, req CreateReportRequest) (*Report, error)
	GetByID(ctx context.Context, id uint) (*Report, error)
	List(ctx context.Context, filter ReportFilter) (types.PagedResult[Report], error)
	Resolve(ctx context.Context, id uint, req ResolveReportRequest) (*Report, error)
	Delete(ctx context.Context, id uint) error
}

type UserRepository interface {
	List(ctx context.Context, filter UserFilter) (types.PagedResult[User], error)
	GetByID(ctx context.Context, userID string) (*User, error)
	BanUser(ctx context.Context, userID string, req BanUserInput) (*User, error)
	UnbanUser(ctx context.Context, userID string, req UnbanUserInput) (*User, error)
}

type AdRepository interface {
	ListAds(ctx context.Context, filter AdFilter) (PagedResult[Ad], error)
	GetAdByID(ctx context.Context, adID string) (*Ad, error)
	CreateAd(ctx context.Context, req CreateAdInput) (*Ad, error)
	UpdateAd(ctx context.Context, adID string, req UpdateAdInput) (*Ad, error)
	SetAdActive(ctx context.Context, adID string, req SetAdActiveInput) (*Ad, error)
	DeleteAd(ctx context.Context, adID string, req DeleteAdInput) error
}

type reportRepositoryImpl struct {
	db *gorm.DB
}

type userRepositoryImpl struct {
	db *gorm.DB
}

type adRepositoryImpl struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepositoryImpl{db: db}
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func NewAdRepository(db *gorm.DB) AdRepository {
	return &adRepositoryImpl{db: db}
}

func (r *reportRepositoryImpl) List(ctx context.Context, filter ReportFilter) (types.PagedResult[Report], error) {
	var items []Report

	query := r.db.WithContext(ctx).Model(&Report{})

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	if filter.UserID != "" {
		query = query.Where("user_id = ?", filter.UserID)
	}

	if filter.PostID != "" {
		query = query.Where("post_id = ?", filter.PostID)
	}

	//Scope filters
	if filter.University != "" {
		query = query.Where("university = ?", filter.University)
	}
	if filter.Level != "" {
		query = query.Where("level = ?", filter.Level)
	}

	if filter.Sort.Field != "" {
		order := filter.Sort.Field
		if filter.Sort.Desc {
			order += "DESC"
		}

		query = query.Order(order)
	}

	var total int64
	query.Count(&total)

	if filter.Limit > 0 {
		offset := (filter.Page - 1) * filter.Limit
		query = query.Limit(filter.Limit).Offset(offset)
	}

	if err := query.Find(&items).Error; err != nil {
		return types.PagedResult[Report]{}, fmt.Errorf("list reports: %w", err)
	}

	return types.PagedResult[Report]{
		Items: items,
		Total: uint64(total),
		Page:  filter.Page,
		Limit: filter.Limit,
	}, nil
}

func (r *reportRepositoryImpl) GetByID(ctx context.Context, id uint) (*Report, error) {
	var report Report

	if err := r.db.WithContext(ctx).First(&report, id).Error; err != nil {
		return nil, fmt.Errorf("get report: %w", err)
	}

	return &report, nil
}

func (r *reportRepositoryImpl) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&Report{}, id).Error; err != nil {
		return fmt.Errorf("delete report: %w", err)
	}

	return nil
}

func (r *reportRepositoryImpl) Resolve(ctx context.Context, id uint, req ResolveReportRequest) (*Report, error) {
	report, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get report: %w", err)
	}

	report.Status = req.Status
	report.Reason = req.ActionMeta.Reason
	report.Note = req.ActionMeta.Note
	report.AdminID = req.ActionMeta.ActedBy

	if err := r.db.WithContext(ctx).Model(report).Updates(report).Error; err != nil {
		return nil, fmt.Errorf("resolve report: %w", err)
	}

	return report, nil
}

func (r *reportRepositoryImpl) Create(ctx context.Context, req CreateReportRequest) (*Report, error) {
	report := &Report{
		Status: "pending",
		UserID: req.UserID,
		PostID: req.PostID,
		Reason: req.Reason,
	}

	if err := r.db.WithContext(ctx).Create(report).Error; err != nil {
		return nil, fmt.Errorf("create report: %w", err)
	}

	return report, nil
}
