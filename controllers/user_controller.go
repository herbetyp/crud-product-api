package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/handlers"
	"github.com/herbetyp/crud-product-api/models"
	"github.com/herbetyp/crud-product-api/services"
)


func CreateUserContoller(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	user.Password = services.SHA256Encoder(user.Password)

	_, err = handlers.CreateUserHandler(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("User %s created", user.Username)})
}
