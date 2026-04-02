package auth

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

//   Auth
//   - POST /auth/register — register with .edu email
//   - POST /auth/login — login, return JWT
//   - POST /auth/verify-email — verify email token
//   - POST /auth/refresh — refresh JWT

type Repository interface {
	FindOrCreateUser(userId string) (*User, error)
	GetUserById(userId string) (*User, error)
	GetAll() ([]User, error)

	Register(req RegisterRequest) error
	Login(req LoginRequest) (*User, error)
	VerifyEmail(req VerifyEmailRequest) (*User, error)
	Refresh(token string) (*User, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) FindOrCreateUser(userId string) (*User, error) {
	user := &User{}
	result := r.db.Where(User{UserID: userId}).FirstOrCreate(user)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find or create user: %w", result.Error)
	}
	return user, nil
}

func (r *repositoryImpl) GetUserById(userId string) (*User, error) {
	user := &User{}
	if err := r.db.Where("user_id = ?", userId).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found %s", userId)
		}
		return nil, fmt.Errorf("get user: %w", err)
	}
	return user, nil
}

func (r *repositoryImpl) GetAll() ([]User, error) {
	var users []User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("finding all users: %w", result.Error)
	}
	return users, nil
}

func (r *repositoryImpl) Register(req RegisterRequest) error {
	otp := &OTPRecord{
		Email:     req.Email,
		Code:      req.OTPCode,
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}
	if err := r.db.Create(otp).Error; err != nil {
		return fmt.Errorf("register: store otp: %w", err)
	}
	return nil
}

func (r *repositoryImpl) Login(req LoginRequest) (*User, error) {
	user := &User{}
	err := r.db.Where("email = ?", req.Email).First(user).Error

	if err != nil {
		return nil, fmt.Errorf("login: user not found")
	}
	return user, nil
}

func (r *repositoryImpl) VerifyEmail(req VerifyEmailRequest) (*User, error) {
	otp := &OTPRecord{}
	err := r.db.Where("email = ? AND code = ?", req.Email, req.OTP).First(otp).Error
	if err != nil {
		return nil, fmt.Errorf("verify: invalid code or email")
	}

	if time.Now().After(otp.ExpiresAt) {
		return nil, fmt.Errorf("verify: code expired")
	}

	user := &User{}
	err = r.db.Where("email = ?", req.Email).First(user).Error
	if err != nil {
		return nil, fmt.Errorf("verify: user not found")
	}

	user.Verified = true
	if err := r.db.Save(user).Error; err != nil {
		return nil, fmt.Errorf("verify: failed to update users table")
	}
	return user, nil
}

func (r *repositoryImpl) Refresh(token string) (*User,error){
	otp := &OTPRecord{}
	err := r.db.Where("code = ?", token).First(otp).Error
	if err != nil {
		return nil, fmt.Errorf("refresh: invalid token")
	}
	if time.Now().After(otp.ExpiresAt) {
		return nil, fmt.Errorf("refresh: token expired")
	}

	user := &User{}
    err = r.db.Where("email = ?", otp.Email).First(user).Error
    if err != nil {
        return nil, fmt.Errorf("refresh: user not found")
    }
    return user, nil
}