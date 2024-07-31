package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/thanhhaudev/openapi-go/app"
	"github.com/thanhhaudev/openapi-go/app/config"
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
	r := mux.NewRouter()
	d := config.NewDatabase()
	h := handler.NewAppHandler(
		handler.NewTenantHandler(d),
		handler.NewUserHandler(d),
	)

	// Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	// Set routes
	app.SetRoutes(r, h)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3000", r))
}
