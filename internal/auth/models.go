package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID string `gorm:"uniqueIndex;not null" json:"userId"`
	Email  string `gorm:"not null" json:"email"`
}

func (User) TableName() string {
	return "users"
}
