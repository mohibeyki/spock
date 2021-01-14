package middleware

import (
	"log"
	"strings"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	"github.com/mohibeyki/spock/model"
	"github.com/mohibeyki/spock/pkg/config"
)

// CORS middleware
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// JWT middleware
func JWT(exceptionsMap map[string]map[string]bool) gin.HandlerFunc {
	config := config.GetConfig()
	return func(c *gin.Context) {
		if value, ok := exceptionsMap[c.Request.URL.String()][c.Request.Method]; ok && value {
			c.Next()
		} else {
			authHeader := c.GetHeader(config.Auth.Header)
			if strings.HasPrefix(authHeader, config.Auth.Prefix) {
				var payload model.Payload
				header, err := jwt.Verify([]byte(authHeader), config.Auth.Algorithm, &payload)
				if err != nil {
					log.Println("ERROR", err)
					c.AbortWithStatusJSON(401, model.ErrResponse{Message: err.Error()})
				} else {
					log.Println("NO ERROR!", header)
					c.Next()
				}
			} else {
				c.AbortWithStatusJSON(401, model.ErrResponse{Message: "missing authorization header!"})
			}
		}
	}
}
