package ads

import (
	"fmt"

	"gorm.io/gorm"
)

//   Ads
//   - GET /ads/banner — get active banner ad (public)

type Repository interface {
	GetActiveBanner(id string) (*Ad, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) GetActiveBanner(id string) (*Ad, error) {
	ad := &Ad{}
	err := r.db.Where("adId = ?", id).First(ad).Error
	if err != nil {
		return nil, fmt.Errorf("ad not found")
	}
	return ad, nil
}
