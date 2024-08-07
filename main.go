package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/thanhhaudev/openapi-go/app/handler"
	_ "github.com/thanhhaudev/openapi-go/docs"
)

func init() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @title Swagger Example API
// @version 1.0
// @description This is a simple Open API example with Go
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Set routes
	router := handler.Router()

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3000", router))
}
