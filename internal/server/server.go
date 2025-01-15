package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	config "github.com/herbetyp/crud-product-api/internal/configs"
	"github.com/herbetyp/crud-product-api/internal/configs/logger"
	"github.com/herbetyp/crud-product-api/internal/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func RunServer() Server {
	APIConf := config.GetConfig().API

	return Server{
		port:   APIConf.Port,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	logger.Info(fmt.Sprintf("Server running at port: %v", s.port))

	log.Fatal(router.Run(":" + s.port))
}
