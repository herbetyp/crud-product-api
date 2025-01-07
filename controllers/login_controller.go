package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/handlers"
	"github.com/herbetyp/crud-product-api/models"
)

func LoginController(ctx *gin.Context) {
	var login models.Login
	err := ctx.BindJSON(&login)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if login.GranType != "client_credentials" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid grant type"})
		return
	}

	token, err := handlers.LoginHandler(&login)
	if err != nil || token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "token_type": "Bearer", "expires_in": 3600})
}
