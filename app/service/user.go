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
		CreateUser(user *model.User) (*model.User, error)
		DeleteUser(user *model.User) error
	}

	userService struct {
		userRepository        repository.UserRepository
		userMessageRepository repository.UserMessageRepository
		logger                *logrus.Logger
	}
)

// DeleteUser deletes a user
func (u userService) DeleteUser(user *model.User) error {
	if err := u.userRepository.Delete(user); err != nil {
		return err
	}

	return nil
}

// CreateUser creates a new user
func (u userService) CreateUser(user *model.User) (*model.User, error) {
	exists, err := u.userRepository.FindByEmail(user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err // only return error if it's not a record not found error
	}

	if exists != nil {
		return nil, appErr.NewUserAlreadyExistsError()
	}

	if err := u.userRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// FindUserByID retrieves a user by ID
func (u userService) FindUserByID(id uint) (*model.User, error) {
	r, err := u.userRepository.FindByID(id)
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
	return u.userRepository.FindAll()
}

// NewUserService creates a new UserService
func NewUserService(
	userRepo repository.UserRepository,
	userMessageRepo repository.UserMessageRepository,
	logger *logrus.Logger,
) UserService {
	return &userService{
		userRepository:        userRepo,
		userMessageRepository: userMessageRepo,
		logger:                logger,
	}
}
