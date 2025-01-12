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
	APIConf := configs.GetConfig().API
	return Server{
		port:   APIConf.Port,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Printf("Server running at port: %v", s.port)

	if !(os.Getenv("GINMODE") == "local") {
		log.Fatal(router.Run(":" + s.port))
	} else {
		log.Fatal(router.Run(":5000"))
	}
}
