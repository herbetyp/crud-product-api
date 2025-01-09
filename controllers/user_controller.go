package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/handlers"
	"github.com/herbetyp/crud-product-api/models"
	"github.com/herbetyp/crud-product-api/services"
)

func CreateUserContoller(ctx *gin.Context) {
	var user models.UserModel

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	user.Password = services.SHA512Encoder(user.Password)

	_, err = handlers.CreateUserHandler(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("User %s created", user.Username)})
}

func UpdateUserPassController(ctx *gin.Context) {
	id := ctx.Param("user_id")
	if id == "" {
		response := models.Response{Message: "Missing user ID"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		response := models.Response{Message: "Invalid user ID. Only integer are allowed"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var user models.UserModel
	err = ctx.BindJSON(&user)
	if err != nil {
		response := models.Response{Message: "Invalid request payload"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user.Password = services.SHA512Encoder(user.Password)

	updatedUserId, err := handlers.UpdateUserPassHandler(userID, user.Password)
	if updatedUserId != 0 && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if updatedUserId == 0 {
		response := models.Response{Message: "User not found"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
