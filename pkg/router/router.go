package router

import (
	"github.com/mohibeyki/spock/controller"
	"github.com/mohibeyki/spock/pkg/middleware"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// Init initializes the router
func Init(db *gorm.DB) *gin.Engine {
	r := gin.New()

	// Middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())

	api := controller.Controller{DB: db}

	apiRouter := r.Group("/api/v1")
	users := apiRouter.Group("/users")
	{
		users.GET("/", api.GetUsers)
		users.GET("/:id", api.GetUser)
		users.POST("/", api.CreateUser)
		users.PUT("/:id", api.UpdateUser)
		users.DELETE("/:id", api.DeleteUser)
	}

	// Protected routes
	// For authorized access, group protected routes using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := apiRouter.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "secure",
	}))

	// /admin/dashboard endpoint is now protected
	authorized.GET("/dashboard", controller.Dashboard)

	return r
}
