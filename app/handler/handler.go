package handler

type (
	// AppHandler interface
	AppHandler interface {
		TenantHandler
		UserHandler
	}

	// TenantHandler interface
	TenantHandler interface {
		// TODO
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
