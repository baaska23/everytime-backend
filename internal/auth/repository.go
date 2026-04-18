package auth

import (
	"gorm.io/gorm"
)

//   Auth
//   - POST /auth/register — register with .edu email
//   - POST /auth/login — login, return JWT
//   - POST /auth/verify-email — verify email token
//   - POST /auth/refresh — refresh JWT

type Repository interface {
	// FindOrCreateUser(userId string) (*User, error)
	// GetUserByEmail(email string) (*User, error)
	// ListAll() ([]User, error)

	// Register(req RegisterRequest) error
	// Login(req LoginRequest) (*User, error)
	// VerifyEmail(req VerifyEmailRequest) (*User, error)
	// CreateRefreshToken(refreshToken *RefreshToken) error
	// FindActiveRefreshToken(tokenHash string) (*RefreshToken, error)
	// RevokeRefreshTokenByHash(tokenHash string) error
	// RevokeAllRefreshTokensByUserID(userID uint) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}