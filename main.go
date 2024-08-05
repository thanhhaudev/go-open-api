package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/thanhhaudev/openapi-go/app"
	"github.com/thanhhaudev/openapi-go/app/config"
	"github.com/thanhhaudev/openapi-go/app/datastore/mysql"
	"github.com/thanhhaudev/openapi-go/app/handler"
	"github.com/thanhhaudev/openapi-go/app/middleware"
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
	router := mux.NewRouter()
	db := config.NewDatabase()
	redis := config.NewRedisStore()
	logger := config.GetLogger()
	tenantRepo := mysql.NewTenantRepository(db.Conn)
	userRepo := mysql.NewUserRepository(db.Conn)

	appHandler := handler.NewAppHandler(
		handler.NewTenantHandler(tenantRepo, logger, redis),
		handler.NewUserHandler(userRepo, logger),
	)

	authMiddleware := middleware.NewAuthMiddleware(tenantRepo, redis.Client).Verify

	// Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	// Set routes
	app.SetRoutes(router, appHandler, authMiddleware)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3000", router))
}
