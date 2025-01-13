package main

import (
	configs "github.com/herbetyp/crud-product-api/config"
	"github.com/herbetyp/crud-product-api/database"
	"github.com/herbetyp/crud-product-api/server"
)

func main() {
	// Initialize App Configurations
	configs.Init()

	// Initialize Database
	database.StartDatabase()

	// Initialize Server
	inicializeServer := server.InitServer()
	inicializeServer.Run()
}
