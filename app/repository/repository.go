package repository

import "github.com/thanhhaudev/openapi-go/app/model"

type (
	TenantRepository interface {
		Find(appKey, appSecret string) (*model.Tenant, error)
		FindByApiKey(apiKey string) (*model.Tenant, error)
	}

	UserRepository interface {
		FindAll() ([]*model.User, error)
		FindByID(id uint) (*model.User, error)
		Create(user *model.User) error
		Update(user *model.User) error
		Delete(id uint) error
	}
)
