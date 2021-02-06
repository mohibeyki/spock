package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mohibeyki/spock/model"
	"github.com/mohibeyki/spock/service"
)

var err error

// GetUser -> [GET] on /users/:email
func (base *Controller) GetUser(c *gin.Context) {
	email := c.Params.ByName("email")

	user, err := service.GetUserByEmail(base.DB, email)
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
		return
	}

	// Fill return data struct
	data := model.Data{
		TotalData:    totalData,
		FilteredData: filteredData,
		Data:         users,
	}

	c.JSON(200, data)
}

// CreateUser -> [POST] on /signup
func (base *Controller) CreateUser(c *gin.Context) {
	user := new(model.User)

	c.ShouldBindJSON(&user)

	u, err := service.GetUserByEmail(base.DB, user.Email)
	if u != nil && err == nil {
		c.JSON(400, model.ErrResponse{Message: "user with the same email exists"})
		return
	}

	hash := service.HashAndSalt(user.Password)
	user.Password = hash

	user, err = service.CreateUser(base.DB, user)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	token := service.GenerateToken(user)

	c.JSON(200, map[string]interface{}{"token": token, "user": user})
}

// Signin -> [POST] on /signin
func (base *Controller) Signin(c *gin.Context) {
	type signInData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	inputUser := new(signInData)
	c.ShouldBindJSON(&inputUser)

	if len(inputUser.Email) == 0 || len(inputUser.Password) == 0 {
		c.JSON(400, model.ErrResponse{Message: "missing email or password!"})
		return
	}

	user, err := service.GetUserByEmail(base.DB, inputUser.Email)
	if user == nil || err != nil {
		c.JSON(404, model.ErrResponse{Message: "user not found!"})
		return
	}

	if !service.ComparePasswords(user.Password, inputUser.Password) {
		c.JSON(404, model.ErrResponse{Message: "user not found!"})
		return
	}

	token := service.GenerateToken(user)
	c.JSON(200, map[string]interface{}{"token": token, "user": user})
}

// UpdateUser -> [PUT] on /users/:email
func (base *Controller) UpdateUser(c *gin.Context) {
	user := base.GetUserFromContext(c)
	targetEmail := c.Params.ByName("email")

	targetUser, err := service.GetUserByEmail(base.DB, targetEmail)
	if err != nil {
		c.AbortWithStatus(404)
	}

	// only admin can update other users
	if user.Role != "admin" && user.Email != targetUser.Email {
		c.AbortWithStatus(403)
	}

	c.ShouldBindJSON(&targetUser)

	// prevent self promotion
	if user.Role != "admin" && targetUser.Role == "admin" {
		c.AbortWithStatus(403)
	}

	targetUser, err = service.UpdateUser(base.DB, targetUser)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(500)
	}

	c.JSON(200, targetUser)
}

// DeleteUser -> [DEL] on /users/:email
func (base *Controller) DeleteUser(c *gin.Context) {
	user := base.GetUserFromContext(c)
	targetEmail := c.Params.ByName("email")

	targetUser, err := service.GetUserByEmail(base.DB, targetEmail)
	if err != nil {
		c.AbortWithStatus(404)
	}

	if user.Role != "admin" && user.Email != targetUser.Email {
		c.AbortWithStatus(403)
	}

	err = service.DeleteUserByEmail(base.DB, targetEmail)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, gin.H{"email#" + targetEmail: "deleted"})
}
