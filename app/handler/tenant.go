package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/thanhhaudev/openapi-go/app/config"
	"github.com/thanhhaudev/openapi-go/app/datastore/mysql"
	"github.com/thanhhaudev/openapi-go/app/service"
	"github.com/thanhhaudev/openapi-go/app/util"
)

type (
	tenantHandler struct {
		tenantService service.TenantService
		logger        *logrus.Logger
	}

	AccessTokenRequest struct {
		ApiKey    string `json:"api_key"`
		ApiSecret string `json:"api_secret"`
	}

	ExchangeTokenRequest struct {
		RefreshToken string `json:"refresh_token"`
	}
)

// GetAccessToken	godoc
// @Summary      	Exchange refresh token for access token
// @Tags         	auth
// @Accept       	json
// @Produce      	json
// @Param			request body ExchangeTokenRequest true "request body"
// @Success      	200  {object} map[string]interface{}
// @Failure      	400  {object} error.AuthError
// @Router       	/api/auth/exchange [post]
func (t tenantHandler) GetAccessToken(w http.ResponseWriter, r *http.Request) {
	p := ExchangeTokenRequest{}
	err := util.Inputs(r, &p)
	if err != nil {
		t.logger.WithError(err).Error("Failed to parse request body")
		util.Response(w, err.Error(), http.StatusBadRequest)

		return
	}

	data, err := t.tenantService.GetAccessToken(r.Context(), p.RefreshToken)
	if err != nil {
		util.Response(w, err, http.StatusBadRequest)
		return
	}

	util.Response(w, map[string]interface{}{"code": http.StatusOK, "data": data}, http.StatusOK)
}

// GetRefreshToken	godoc
// @Summary      	Get refresh token by API key and secret
// @Tags         	auth
// @Accept       	json
// @Produce      	json
// @Param			request body AccessTokenRequest true "request body"
// @Success      	200  {object} map[string]interface{}
// @Failure      	400  {object} error.AuthError
// @Router       	/api/auth/access [post]
func (t tenantHandler) GetRefreshToken(w http.ResponseWriter, r *http.Request) {
	p := AccessTokenRequest{}
	err := util.Inputs(r, &p)
	if err != nil {
		t.logger.WithError(err).Error("Failed to parse request body")
		util.Response(w, err.Error(), http.StatusBadRequest)

		return
	}

	data, err := t.tenantService.GetRefreshToken(r.Context(), p.ApiKey, p.ApiSecret)
	if err != nil {
		t.logger.WithError(err).Error("Failed to get refresh token")
		util.Response(w, err, http.StatusBadRequest)

		return
	}

	util.Response(w, map[string]interface{}{"code": http.StatusOK, "data": data}, http.StatusOK)
}

// NewTenantHandler creates a new TenantHandler
func NewTenantHandler(db *config.Database, s *config.RedisStore) TenantHandler {
	r := mysql.NewTenantRepository(db.Conn)
	l := config.GetLogger()

	return &tenantHandler{
		tenantService: service.NewTenantService(r, s.Client, l),
		logger:        l,
	}
}
