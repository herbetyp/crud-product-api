package main

import (
	"github.com/herbetyp/crud-product-api/internal/configs/logger"

	config "github.com/herbetyp/crud-product-api/internal/configs"
	"github.com/herbetyp/crud-product-api/internal/database"
	"github.com/herbetyp/crud-product-api/internal/server"
)

func main() {
	logger.Info("Starting application...")

	// Loading App Config
	config.InitConfig()

	// Connecting on Database
	database.StartDatabase()

	// Starting Server
	runServer := server.RunServer()
	runServer.Run()
}
