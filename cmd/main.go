package main

import (
	"github.com/herbetyp/crud-product-api/configs"
	"github.com/herbetyp/crud-product-api/database"
	"github.com/herbetyp/crud-product-api/server"
)

func main() {
	// Initialize database
	configs.GetConfig()
	database.ConnectDB()
	
	// Initialize the server
	inicializeServer := server.InitServer()
	inicializeServer.Run()
}
