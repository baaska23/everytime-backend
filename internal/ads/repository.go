package ads

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

//   Ads
//   - GET /ads/banner — get active banner ad (public)

type AdRepository interface {
	GetActiveBanner(ctx context.Context, id string) (*Ad, error)
	ListActiveBanners(ctx context.Context) ([]Ad, error)
}

type adRepositoryImpl struct {
	db *gorm.DB
}

func NewAdRepository(db *gorm.DB) AdRepository {
	return &adRepositoryImpl{db: db}
}

func (r *adRepositoryImpl) GetActiveBanner(ctx context.Context, id string) (*Ad, error) {
	ad := &Ad{}
	err := r.db.WithContext(ctx).First(ad, id).Error
	if err != nil {
		return nil, fmt.Errorf("ad not found")
	}
	return ad, nil
}

func (r *adRepositoryImpl) ListActiveBanners(ctx context.Context) ([]Ad, error) {
	var ads []Ad
	err := r.db.WithContext(ctx).Where("end_date >= ?", time.Now()).Find(&ads).Error
	if err != nil {
		return nil, fmt.Errorf("Cannot list active ads")
	}
	return ads, nil
}
