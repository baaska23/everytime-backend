package admin

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	PostID string `json:"postId"`
	UserID string `json:"userId"`
	Status string `gorm:"default:pending" json:"status"`
}

type User struct {
	gorm.Model
	Email      string `json:"email"`
	University string `json:"university"`
	Faculty    string `json:"faculty"`
	Department string `json:"department"`
}

type Ad struct {
	gorm.Model
	BannerUrl string    `gorm:"not null" json:"bannerUrl"`
	StartDate time.Time `gorm:"not null" json:"startDate"`
	EndDate   time.Time `gorm:"not null" json:"endDate"`
}

type ResolveReportRequest struct {
	Status string `json:"status"`
}
