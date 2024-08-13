package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/thanhhaudev/openapi-go/app/common"
	"github.com/thanhhaudev/openapi-go/app/config"
	"github.com/thanhhaudev/openapi-go/app/datastore/mysql"
	"github.com/thanhhaudev/openapi-go/app/repository"
	"github.com/thanhhaudev/openapi-go/app/util"
)

var (
	routeHandler    AppHandler
	routeMap        map[string][]*route // map[scope][route]
	db              *config.Database
	redisStore      *config.RedisStore
	logger          *logrus.Logger
	tenantRepo      repository.TenantRepository
	userRepo        repository.UserRepository
	userMessageRepo repository.UserMessageRepository
	messageRepo     repository.MessageRepository
)

// inject dependencies
func inject() {
	db = config.NewDatabase()
	redisStore = config.NewRedisStore()
	logger = config.GetLogger()
	tenantRepo = mysql.NewTenantRepository(db.Conn)
	userRepo = mysql.NewUserRepository(db.Conn)
	messageRepo = mysql.NewMessageRepository(db.Conn)
	userMessageRepo = mysql.NewUserMessageRepository(db.Conn)
	routeHandler = NewAppHandler(
		NewTenantHandler(tenantRepo, logger, redisStore),
		NewUserHandler(userRepo, userMessageRepo, logger),
		NewMessageHandler(userRepo, userMessageRepo, messageRepo, logger),
	)

	routeMap = map[string][]*route{
		common.ScopeManageUser: {
			{
				routeHandler.CreateUser,
				"/api/v1/users",
				http.MethodPost,
			},
			{
				routeHandler.GetUsers,
				"/api/v1/users",
				http.MethodGet,
			},
			{
				routeHandler.FindUser,
				"/api/v1/users/{id:[0-9]+}",
				http.MethodGet,
			},
			{
				routeHandler.DeleteUser,
				"/api/v1/users/{id:[0-9]+}",
				http.MethodDelete,
			},
			{
				routeHandler.UpdateUser,
				"/api/v1/users/{id:[0-9]+}",
				http.MethodPut,
			},
			{
				routeHandler.UserMessages,
				"/api/v1/users/{id:[0-9]+}/messages",
				http.MethodGet,
			},
		},
		common.ScopeManageMessage: {
			{
				routeHandler.CreateMessage,
				"/api/v1/messages",
				http.MethodPost,
			},
		},
	}
}

// Router godoc
func Router() *mux.Router {
	router := mux.NewRouter()
	inject()
	setupSwagger(router)

	router.HandleFunc("/health", func(w http.ResponseWriter, request *http.Request) {
		util.Response(w, map[string]bool{"ok": true}, http.StatusOK)
	})

	router.HandleFunc("/api/auth/access", routeHandler.GetRefreshToken).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/exchange", routeHandler.GetAccessToken).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/refresh", routeHandler.RefreshAccessToken).Methods(http.MethodPost)

	// Authenticated routes
	r := router.NewRoute().Subrouter()
	r.Use(verifyToken)
	r.Use(verifyScope)

	// apply routes
	for _, routes := range routeMap {
		for _, route := range routes {
			r.HandleFunc(route.Path, route.HandlerFunc).Methods(route.Method)
		}
	}

	return router
}

// setupSwagger godoc
func setupSwagger(r *mux.Router) {
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
}
