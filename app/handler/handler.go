package handler

import (
	"net/http"
)

type (
	// AppHandler interface
	AppHandler interface {
		TenantHandler
		UserHandler
		MessageHandler
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
		GetUser(http.ResponseWriter, *http.Request)
		CreateUser(http.ResponseWriter, *http.Request)
		UpdateUser(http.ResponseWriter, *http.Request)
		DeleteUser(http.ResponseWriter, *http.Request)
		GetUserMessages(http.ResponseWriter, *http.Request)
	}

	MessageHandler interface {
		CreateMessage(http.ResponseWriter, *http.Request)
		GetMessage(http.ResponseWriter, *http.Request)
	}

	// Handler struct
	appHandler struct {
		TenantHandler
		UserHandler
		MessageHandler
	}

	route struct {
		HandlerFunc http.HandlerFunc
		Path        string
		Method      string
	}
)

// NewAppHandler godoc
func NewAppHandler(t TenantHandler, u UserHandler, m MessageHandler) AppHandler {
	return &appHandler{
		TenantHandler:  t,
		UserHandler:    u,
		MessageHandler: m,
	}
}
