package main

import (
	"github.com/herbetyp/crud-product-api/server"
)

func main() {
	// Initialize the server
	inicializeServer := server.InitServer()
	inicializeServer.Run()
}
