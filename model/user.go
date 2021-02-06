package model

import (
	"gorm.io/gorm"
)

// User is the basic user model
type User struct {
	gorm.Model `json:"-"`
	Email      string `json:"email" gorm:"index"`
	Role       string `json:"role"`
	Avatar     string `json:"avatar"`
	Bio        string `json:"bio" gorm:"type:text"`
	Password   string `json:"-"`
}
