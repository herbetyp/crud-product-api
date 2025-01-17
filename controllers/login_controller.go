package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/handlers"
	config "github.com/herbetyp/crud-product-api/internal/configs"
	"github.com/herbetyp/crud-product-api/internal/configs/logger"
	model "github.com/herbetyp/crud-product-api/models/login"
	repository "github.com/herbetyp/crud-product-api/repositories"
	zapLog "go.uber.org/zap"
)

func Login(c *gin.Context) {
	var dto model.LoginDTO

	err := c.BindJSON(&dto)
	if err != nil {
		logger.Error("Invalid payload: ", err)
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if dto.GranType != "client_credentials" {
		var err error
		logger.Error("Invalid grant type: ", err, zapLog.String("grant_type", dto.GranType))
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid grant type"})
		return
	}

	repo := &repository.LoginRepository{}
	handler := handlers.NewLoginHandler(repo)

	token, tokenId, err := handler.NewLogin(dto)

	if err != nil || token == "" {
		logger.Error("Error on login", err)
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// Log
	logger.Info("Token issued",
		zapLog.String("request_id", c.Param("request_id")),
		zapLog.String("ip", c.Param("ip")),
		zapLog.String("email", dto.Email),
		zapLog.String("jwt_id", tokenId),
		zapLog.String("method", c.Param("method")),
		zapLog.String("path", c.Param("path")),
		zapLog.String("protocol", c.Param("protocol")),
		zapLog.String("user_agent", c.Param("user_agent")),
	)

	JWTConf := config.GetConfig().JWT
	c.JSON(200, gin.H{"access_token": token, "token_type": "Bearer", "expires_in": JWTConf.ExpiresIn})
}
