package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/thanhhaudev/openapi-go/app/common"
	"github.com/thanhhaudev/openapi-go/app/config"
	"github.com/thanhhaudev/openapi-go/app/datastore/mysql"
	"github.com/thanhhaudev/openapi-go/app/middleware"
	"github.com/thanhhaudev/openapi-go/app/util"
)

var (
	routeHandler    AppHandler
	routeMap        map[string][]*route // map[scope][route]
	routeMiddleware []func(next http.Handler) http.Handler
)

func inject() {
	db := config.NewDatabase()
	redisStore := config.NewRedisStore()
	logger := config.GetLogger()
	tenantRepo := mysql.NewTenantRepository(db.Conn)
	userRepo := mysql.NewUserRepository(db.Conn)

	authMiddleware := middleware.NewAuthMiddleware(tenantRepo, redisStore.Client)
	routeMiddleware = []func(next http.Handler) http.Handler{
		authMiddleware.VerifyToken,
	}

	routeHandler = NewAppHandler(
		NewTenantHandler(tenantRepo, logger, redisStore),
		NewUserHandler(userRepo, logger),
	)

	routeMap = map[string][]*route{
		common.ScopeManageUser: {
			{
				routeHandler.GetUsers,
				"/users",
				http.MethodGet,
			},
			{
				routeHandler.FindUser,
				"/users/{id:[0-9]+}",
				http.MethodGet,
			},
		},
	}
}

// Router godoc
func Router() *mux.Router {
	router := mux.NewRouter()
	inject()
	setupSwagger(router)

	s := router.PathPrefix("/api").Subrouter()
	s.HandleFunc("/health", func(w http.ResponseWriter, request *http.Request) {
		util.Response(w, map[string]bool{"ok": true}, http.StatusOK)
	})

	a := s.PathPrefix("/auth").Subrouter()
	a.HandleFunc("/access", routeHandler.GetRefreshToken).Methods(http.MethodPost)
	a.HandleFunc("/exchange", routeHandler.GetAccessToken).Methods(http.MethodPost)
	a.HandleFunc("/refresh", routeHandler.RefreshAccessToken).Methods(http.MethodPost)

	// Authenticated routes
	r := s.PathPrefix("/v1").Subrouter()
	for _, mw := range routeMiddleware {
		r.Use(mw)
	}

	// apply routes
	for _, routes := range routeMap {
		for _, route := range routes {
			r.HandleFunc(route.Path, route.HandlerFunc).Methods(route.Method)
		}
	}

	return router
}

// SetupSwagger godoc
func setupSwagger(r *mux.Router) {
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
}
