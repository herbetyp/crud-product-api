package main

import (
	"github.com/herbetyp/crud-product-api/server"
)

func main() {
	// Initialize application
	inicializeServer := server.InitServer()
	inicializeServer.Run()
}
