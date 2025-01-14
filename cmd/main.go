package main

import (
	"github.com/herbetyp/crud-product-api/config/logger"

	"github.com/herbetyp/crud-product-api/config"
	"github.com/herbetyp/crud-product-api/database"
	"github.com/herbetyp/crud-product-api/server"
)

func main() {
	logger.Info("starting application...")

	// Loading App Config
	config.InitConfig()

	// Connecting on Database
	database.StartDatabase()

	// Starting Server
	runServer := server.RunServer()
	runServer.Run()
}
