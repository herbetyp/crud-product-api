package main

import (
	"github.com/herbetyp/crud-product-api/configs"
	"github.com/herbetyp/crud-product-api/database"
	"github.com/herbetyp/crud-product-api/server"
)

func main() {
	configs.GetConfig()
	database.ConnectDB()

	// Initialize application
	inicializeServer := server.InitServer()
	inicializeServer.Run()
}
