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
		GetUser() string
	}

	// Handler struct
	appHandler struct {
		TenantHandler
		UserHandler
	}
)

// NewAppHandler godoc
func NewAppHandler(t TenantHandler, u UserHandler) AppHandler {
	return &appHandler{
		TenantHandler: t,
		UserHandler:   u,
	}
}
