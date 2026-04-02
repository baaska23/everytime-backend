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
	Verified   bool   `gorm:"default:false" json:"verified"`
}

func (User) TableName() string { return "users" }

type OTPRecord struct {
	gorm.Model
	Email     string    `gorm:"not null;index"`
	Code      string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	Used      bool      `gorm:"default:false"`
}

func (OTPRecord) TableName() string { return "otp_records" }

type RegisterRequest struct {
	Email   string `json:"email"`
	OTPCode string `json:"-"` // set by service, not from HTTP body
}

type VerifyEmailRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type LoginRequest struct {
	Email string `json:"email"`
}
