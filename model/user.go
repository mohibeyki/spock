package model

import (
	"gorm.io/gorm"
)

// User is the basic user model
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"index"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio" gorm:"type:text"`
}
