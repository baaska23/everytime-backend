package ads

import (
	"time"

	"gorm.io/gorm"
)

type Ad struct {
	gorm.Model
	BannerUrl string    `gorm:"not null" json:"bannerUrl"`
	StartDate time.Time `gorm:"not null" json:"startDate"`
	EndDate   time.Time `gorm:"not null" json:"endDate"`
}
