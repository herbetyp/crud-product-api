package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/handlers"
)

func LoginController(ctx *gin.Context) {
	handlers.LoginHandler()

	ctx.JSON(http.StatusOK, gin.H{"message": "Login success"})
}
