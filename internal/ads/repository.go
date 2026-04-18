package ads

import (
	"context"

	"gorm.io/gorm"
)

type AdRepository interface {
	ListActiveBanners(ctx context.Context) ([]Ad, error)
}

type adRepositoryImpl struct {
	db *gorm.DB
}

func NewAdRepository(db *gorm.DB) AdRepository {
	return &adRepositoryImpl{db: db}
}
