package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/herbetyp/crud-product-api/handlers"
	model "github.com/herbetyp/crud-product-api/models/user"
	repository "github.com/herbetyp/crud-product-api/repositories"
	service "github.com/herbetyp/crud-product-api/services"
)

func CreateUser(c *gin.Context) {
	var dto model.UserDTO

	err := c.BindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if !service.ValidateEmailFormat(dto.Email) {
		c.JSON(400, gin.H{"error": "Invalid email format"})
		return
	}

	repo := &repository.UserRepository{}
	handler := handlers.NewUserHandler(repo)

	result, err := handler.CreateUser(dto)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(201, result)
}

func GetUser(c *gin.Context) {
	id := c.Param("user_id")
	if id == "" {
		c.JSON(400, "Missing user id")
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	repo := &repository.UserRepository{}
	handler := handlers.NewUserHandler(repo)

	result, err := handler.GetUser(uint(userId))

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func GetUsers(c *gin.Context) {
	repo := &repository.UserRepository{}
	handler := handlers.NewUserHandler(repo)

	result, err := handler.GetUsers()

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("user_id")

	if id == "" {
		c.JSON(400, "Missing user ID")
		return
	}

	userID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var dto model.UserDTO

	err = c.BindJSON(&dto)
	if err != nil {
		c.JSON(400, "Invalid request payload")
		return
	}

	repo := &repository.UserRepository{}
	handler := handlers.NewUserHandler(repo)

	dto.ID = uint(userID)
	result, err := handler.UpdateUser(dto)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("user_id")
	if id == "" {
		c.JSON(400, "Missing user ID")
		return
	}

	userID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var dto model.UserDTO

	repo := &repository.UserRepository{}
	handler := handlers.NewUserHandler(repo)

	dto.ID = uint(userID)
	result, err := handler.DeleteUser(dto)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}


func RecoveryUser(c *gin.Context) {
	id := c.Param("user_id")
	if id == "" {
		c.JSON(400, "Missing user ID")
		return
	}

	userID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var dto model.UserDTO

	repo := &repository.UserRepository{}
	handler := handlers.NewUserHandler(repo)

	dto.ID = uint(userID)
	result, err := handler.RecoveryUser(dto)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}