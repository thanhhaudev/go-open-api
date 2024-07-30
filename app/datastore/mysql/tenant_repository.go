package mysql

import (
	"github.com/thanhhaudev/openapi-go/app/model"
	"github.com/thanhhaudev/openapi-go/app/repository"
	"gorm.io/gorm"
)

type tenantRepository struct {
	gorm *gorm.DB
}

// FindByKey finds a tenant by app key and app secret
func (t tenantRepository) FindByKey(appKey, appSecret string) (*model.Tenant, error) {
	tenant := &model.Tenant{}

	err := t.gorm.
		Where(`app_key = ? AND app_secret = ?`, appKey, appSecret).
		First(tenant).Error
	if err != nil {
		return nil, err
	}

	return tenant, nil
}

// NewTenantRepository creates a new tenant repository
func NewTenantRepository(gorm *gorm.DB) repository.TenantRepository {
	return tenantRepository{
		gorm: gorm,
	}
}
