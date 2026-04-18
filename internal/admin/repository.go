package admin

import (
	"context"

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

type ReportRepository interface {
    CreateReport(ctx context.Context, req CreateReportInput) (*Report, error)
    GetReportByID(ctx context.Context, reportID string) (*Report, error)
    ListReports(ctx context.Context, filter ReportFilter) (PagedResult[Report], error)
    ResolveReport(ctx context.Context, reportID string, req ResolveReportInput) (*Report, error)
    DeleteReport(ctx context.Context, reportID string) error
}

type UserRepository interface {
    ListUsers(ctx context.Context, filter UserFilter) (PagedResult[User], error)
    GetUserByID(ctx context.Context, userID string) (*User, error)
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

type Pagination struct {
	Page  int
	Limit int
}

type SortOption struct {
	Field string
	Desc  bool
}

type Scope struct {
	University string
}

type ActionMeta struct {
	ActedBy string
	Reason  string
	Note    string
}

type PagedResult[T any] struct {
	Items []T
	Total int64
	Page  int
	Limit int
}

type ReportFilter struct {
	Pagination
	Scope
	Status   string
	PostID   string
	UserID   string
	FromDate string
	ToDate   string
	Sort     SortOption
}

type UserFilter struct {
	Pagination
	Scope
	Search     string
	IsBanned   *bool
	Department string
	Faculty    string
	Sort       SortOption
}

type AdFilter struct {
	Pagination
	Scope
	IsActive *bool
	Slot     string
	Sort     SortOption
}

type CreateReportInput struct {
	PostID string
	UserID string
	Reason string
	Scope  Scope
}

type ResolveReportInput struct {
	Status string
	ActionMeta
}

type BulkResolveReportsInput struct {
	ReportIDs []string
	Status    string
	ActionMeta
}

type BanUserInput struct {
	Until string
	ActionMeta
}

type UnbanUserInput struct {
	ActionMeta
}

type CreateAdInput struct {
	BannerURL string
	StartDate string
	EndDate   string
	Priority  int
	Slot      string
	Scope     Scope
	ActionMeta
}

type UpdateAdInput struct {
	BannerURL *string
	StartDate *string
	EndDate   *string
	Priority  *int
	Slot      *string
	ActionMeta
}

type SetAdActiveInput struct {
	IsActive bool
	ActionMeta
}

type DeleteAdInput struct {
	HardDelete bool
	ActionMeta
}

type BulkSetAdActiveInput struct {
	AdIDs    []string
	IsActive bool
	ActionMeta
}

type adminRepositoryImpl struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepositoryImpl{db: db}
}
