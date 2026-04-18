package auth

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID     string `gorm:"uniqueIndex;not null" json:"userId"`
	Email      string `gorm:"not null" json:"email"`
	University string `gorm:"not null" json:"university"`
	Faculty    string `gorm:"not null" json:"faculty"`
	Department string `json:"department"`
	Verified   bool   `gorm:"default:false" json:"verified"`
}

type OTPRecord struct {
	gorm.Model
	Email     string    `gorm:"not null;index"`
	Code      string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	Used      bool      `gorm:"default:false"`
}

type RefreshToken struct {
	gorm.Model
	UserID    uint       `gorm:"not null;index"`
	TokenHash string     `gorm:"size:64;uniqueIndex;not null"`
	ExpiresAt time.Time  `gorm:"not null;index"`
	RevokedAt *time.Time `gorm:"index"`
}

type RegisterRequest struct {
	Email   string `json:"email"`
	OTPCode string `json:"-"`
}

type VerifyEmailRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type LoginRequest struct {
	Email string `json:"email"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
