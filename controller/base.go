package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mohibeyki/spock/model"
	"gorm.io/gorm"
)

// Controller is a generic controller struct
type Controller struct {
	DB *gorm.DB
}

// GetUserFromContext extracts the user form context
func (base *Controller) GetUserFromContext(c *gin.Context) *model.User {
	tmp, _ := c.Get("user")
	return tmp.(*model.User)
}
