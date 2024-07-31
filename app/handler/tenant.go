package handler

import (
	"github.com/thanhhaudev/openapi-go/app/config"
	"github.com/thanhhaudev/openapi-go/app/datastore/mysql"
	"github.com/thanhhaudev/openapi-go/app/service"
)

type tenantHandler struct {
	tenantService service.TenantService
}

func NewTenantHandler(db *config.Database) TenantHandler {
	r := mysql.NewTenantRepository(db.Conn)

	return &tenantHandler{
		tenantService: service.NewTenantService(r),
	}
}
