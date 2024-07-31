package service

import "github.com/thanhhaudev/openapi-go/app/repository"

type (
	TenantService interface {
		// todo
	}

	tenantService struct {
		TenantRepository repository.TenantRepository
	}
)

// NewTenantService creates a new TenantService
func NewTenantService(r repository.TenantRepository) TenantService {
	return &tenantService{
		TenantRepository: r,
	}
}
