package controller

import "gorm.io/gorm"

// Controller is a generic controller struct
type Controller struct {
	DB *gorm.DB
}
