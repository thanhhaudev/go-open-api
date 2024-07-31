package handler

import (
	"net/http"

	"github.com/thanhhaudev/openapi-go/app/config"
	"github.com/thanhhaudev/openapi-go/app/datastore/mysql"
	"github.com/thanhhaudev/openapi-go/app/service"
	"github.com/thanhhaudev/openapi-go/app/util"
)

type (
	tenantHandler struct {
		tenantService service.TenantService
	}

	AccessTokenRequest struct {
		ApiKey    string `json:"api_key"`
		ApiSecret string `json:"api_secret"`
	}
)

// GetRefreshToken godoc
func (t tenantHandler) GetRefreshToken(w http.ResponseWriter, r *http.Request) {
	p := AccessTokenRequest{}
	err := util.Inputs(r, &p)
	if err != nil {
		util.Response(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := t.tenantService.GetRefreshToken(p.ApiKey, p.ApiSecret)
	if err != nil {
		util.Response(w, err, http.StatusBadRequest)
		return
	}

	util.Response(w, map[string]interface{}{"code": http.StatusOK, "data": data}, http.StatusOK)
}

// NewTenantHandler creates a new TenantHandler
func NewTenantHandler(db *config.Database) TenantHandler {
	r := mysql.NewTenantRepository(db.Conn)

	return &tenantHandler{
		tenantService: service.NewTenantService(r),
	}
}
