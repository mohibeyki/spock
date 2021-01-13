package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dashboard is a sample controller that uses basic http auth
func Dashboard(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	c.JSON(http.StatusOK, gin.H{"Welcome: ": user})
}
