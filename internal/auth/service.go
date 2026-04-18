package auth

import (
	"os"
	"time"
)

const (
	defaultAccessTokenTTL  = 15 * time.Minute
	defaultRefreshTokenTTL = 30 * 24 * time.Hour
)

type Service struct {
	repo       Repository
	jwtSecret  []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewService(repo Repository) *Service {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "dev-secret-change-me"
	}

	return &Service{
		repo:       repo,
		jwtSecret:  []byte(secret),
		accessTTL:  defaultAccessTokenTTL,
		refreshTTL: defaultRefreshTokenTTL,
	}
}