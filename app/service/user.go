package service

import "github.com/thanhhaudev/openapi-go/app/repository"

type (
	UserService interface {
		// todo
	}

	userService struct {
		UserRepository repository.UserRepository
	}
)

// NewUserService creates a new UserService
func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		UserRepository: r,
	}
}
