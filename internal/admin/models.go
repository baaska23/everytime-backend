package admin

import (
	"everytime-backend/internal/shared/types"
	"time"

	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	PostID  string `json:"postId"`
	UserID  string `json:"userId"`
	AdminID string `json:"adminId"`
	Status  string `gorm:"default:pending" json:"status"`
	Reason  string `json:"reason"`
	Note    string `json:"note"`
}

type User struct {
	gorm.Model
	Role       string `json:"role"`
	Email      string `json:"email"`
	University string `json:"university"`
	Faculty    string `json:"faculty"`
	Department string `json:"department"`
	Level      string `json:"level"`
}

type Ad struct {
	gorm.Model
	BannerUrl string    `gorm:"not null" json:"bannerUrl"`
	StartDate time.Time `gorm:"not null" json:"startDate"`
	EndDate   time.Time `gorm:"not null" json:"endDate"`
}

type CreateReportRequest struct {
	PostID string
	UserID string
	Reason string
}

type ResolveReportRequest struct {
	Status string
	types.ActionMeta
}

type BanUserInput struct {
	Until string
	types.ActionMeta
}

type UnbanUserInput struct {
	types.ActionMeta
}

type CreateAdInput struct {
	BannerURL string
	StartDate string
	EndDate   string
	Priority  int
	Slot      string
	Scope     types.Scope
	types.ActionMeta
}

type UpdateAdInput struct {
	BannerURL *string
	StartDate *string
	EndDate   *string
	Priority  *int
	Slot      *string
	types.ActionMeta
}

type ReportFilter struct {
	types.Pagination
	types.Scope
	Status string
	UserID string
	PostID string
	Sort   types.SortOption
}
