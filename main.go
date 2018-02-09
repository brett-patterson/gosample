package main

import (
	"github.com/brett-patterson/gosample/comm"
)

// Main entry point to application
func main() {
	server := comm.CreateServer()
	server.Logger.Fatal(server.Start(":8080"))
}
