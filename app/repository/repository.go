package repository

import "github.com/thanhhaudev/openapi-go/app/model"

type (
	TenantRepository interface {
		FindByKeys(appKey, appSecret string) (*model.Tenant, error)
	}

	UserRepository interface {
		FindAll() ([]*model.User, error)
		FindByID(id int64) (*model.User, error)
		Create(user *model.User) error
		Update(user *model.User) error
		Delete(id int64) error
	}
)
