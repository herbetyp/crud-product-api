package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/crud-product-api/configs"
	"github.com/herbetyp/crud-product-api/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func InitServer() Server {
	cfg := configs.GetConfig()
	return Server{
		port:   cfg.API.Port,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Printf("Server running at port: %v", s.port)

	if !(os.Getenv("MODE") == "DEBUG") {
		log.Fatal(router.Run(":" + s.port))
	} else {
		log.Fatal(router.Run(":5000"))
	}
}
