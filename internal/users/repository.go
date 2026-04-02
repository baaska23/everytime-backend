package users

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindOrCreateUser(userId string) (*User, error)
	GetUserById(userId string) (*User, error)
	GetAll() ([]User, error)
}

type everytimeRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &everytimeRepository{db: db}
}

func (r *everytimeRepository) FindOrCreateUser(userId string) (*User, error) {
	user := &User{}
	result := r.db.Where(User{UserID: userId}).FirstOrCreate(user)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find or create user: %w", result.Error)
	}
	return user, nil
}

func (r *everytimeRepository) GetUserById(userId string) (*User, error) {
	user := &User{}
	if err := r.db.Where("user_id = ?", userId).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found %s", userId)
		}
		return nil, fmt.Errorf("get user: %w", err)
	}
	return user, nil
}

func (r *everytimeRepository) GetAll() ([]User, error) {
	var users []User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("finding all users: %w", result.Error)
	}
	return users, nil
}
