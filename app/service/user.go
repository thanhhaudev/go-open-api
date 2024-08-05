package service

import (
	"errors"

	"github.com/sirupsen/logrus"
	appErr "github.com/thanhhaudev/openapi-go/app/error"
	"github.com/thanhhaudev/openapi-go/app/model"
	"github.com/thanhhaudev/openapi-go/app/repository"
	"gorm.io/gorm"
)

type (
	UserService interface {
		GetUsers() ([]*model.User, error)
		FindUserByID(id uint) (*model.User, error)
	}

	userService struct {
		UserRepository repository.UserRepository
		logger         *logrus.Logger
	}
)

// FindUserByID retrieves a user by ID
func (u userService) FindUserByID(id uint) (*model.User, error) {
	r, err := u.UserRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErr.NewUserNotFoundError()
		}

		return nil, err
	}

	return r, nil
}

// GetUsers retrieves all users
func (u userService) GetUsers() ([]*model.User, error) {
	return u.UserRepository.FindAll()
}

// NewUserService creates a new UserService
func NewUserService(r repository.UserRepository, l *logrus.Logger) UserService {
	return &userService{
		UserRepository: r,
		logger:         l,
	}
}
