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

func main() {
	// Set routes
	router := handler.Router()

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3000", router))
}
