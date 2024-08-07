package handler

import (
	"net/http"
)

type (
	// AppHandler interface
	AppHandler interface {
		TenantHandler
		UserHandler
	}

	// TenantHandler interface
	TenantHandler interface {
		GetRefreshToken(http.ResponseWriter, *http.Request)
		GetAccessToken(http.ResponseWriter, *http.Request)
		RefreshAccessToken(http.ResponseWriter, *http.Request)
	}

	// UserHandler interface
	UserHandler interface {
		GetUsers(http.ResponseWriter, *http.Request)
		FindUser(http.ResponseWriter, *http.Request)
		CreateUser(http.ResponseWriter, *http.Request)
		UpdateUser(http.ResponseWriter, *http.Request)
		DeleteUser(http.ResponseWriter, *http.Request)
		UserMessages(http.ResponseWriter, *http.Request)
	}

	MessageHandler interface {
		// TODO
	}

	// Handler struct
	appHandler struct {
		TenantHandler
		UserHandler
	}

	route struct {
		HandlerFunc http.HandlerFunc
		Path        string
		Method      string
	}
)

// NewAppHandler godoc
func NewAppHandler(t TenantHandler, u UserHandler) AppHandler {
	return &appHandler{
		TenantHandler: t,
		UserHandler:   u,
	}
}
