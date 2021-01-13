package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mohibeyki/spock/model"
	"github.com/mohibeyki/spock/service"
)

var err error

// GetUser -> [GET] on /users/:id
func (base *Controller) GetUser(c *gin.Context) {
	id := c.Params.ByName("id")

	user, err := service.GetUser(base.DB, id)
	if err != nil {
		c.AbortWithStatus(404)
	}

	c.JSON(200, user)
}

// GetUsers -> [GET] on /users
func (base *Controller) GetUsers(c *gin.Context) {
	var args model.Args

	// Define and get sorting field
	args.Sort = c.DefaultQuery("Sort", "ID")

	// Define and get sorting order field
	args.Order = c.DefaultQuery("Order", "DESC")

	// Define and get offset for pagination
	args.Offset = c.DefaultQuery("Offset", "0")

	// Define and get limit for pagination
	args.Limit = c.DefaultQuery("Limit", "20")

	// Get search keyword for Search Scope
	args.Search = c.DefaultQuery("Search", "")

	// Fetch results from database
	users, filteredData, totalData, err := service.GetUsers(c, base.DB, args)
	if err != nil {
		c.AbortWithStatus(404)
	}

	// Fill return data struct
	data := model.Data{
		TotalData:    totalData,
		FilteredData: filteredData,
		Data:         users,
	}

	c.JSON(200, data)
}

// CreateUser -> [POST] on /users/
func (base *Controller) CreateUser(c *gin.Context) {
	user := new(model.User)

	c.ShouldBindJSON(&user)

	user, err := service.CreateUser(base.DB, user)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, user)
}

// UpdateUser -> [PUT] on /users/:id
func (base *Controller) UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")

	user, err := service.GetUser(base.DB, id)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.ShouldBindJSON(&user)

	user, err = service.UpdateUser(base.DB, user)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, user)
}

// DeleteUser -> [DEL] on /users/:id
func (base *Controller) DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")

	err = service.DeleteUser(base.DB, id)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, gin.H{"id#" + id: "deleted"})
}
