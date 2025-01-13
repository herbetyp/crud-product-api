package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/handlers"
	model "github.com/herbetyp/crud-product-api/models/login"
	repository "github.com/herbetyp/crud-product-api/repositories"
	service "github.com/herbetyp/crud-product-api/services"
)

func Login(c *gin.Context) {
	var dto model.LoginDTO

	err := c.BindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if dto.GranType != "client_credentials" {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid grant type"})
		return
	}

	repo := &repository.LoginRepository{}
	handler := handlers.NewLoginHandler(repo)

	token, err := handler.NewLogin(dto)
	
	if err != nil || token == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(200, gin.H{"token": token, "token_type": "Bearer", "expires_in": 3600})
}
