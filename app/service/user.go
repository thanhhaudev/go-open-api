package service

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/thanhhaudev/go-open-api/app/command"
	appErr "github.com/thanhhaudev/go-open-api/app/error"
	"github.com/thanhhaudev/go-open-api/app/model"
	"github.com/thanhhaudev/go-open-api/app/repository"
	"gorm.io/gorm"
)

type (
	UserService interface {
		GetUsers() ([]*model.User, error)
		FindUserByID(id uint) (*model.User, error)
		CreateUser(user *model.User) (*model.User, error)
		UpdateUser(id uint, com *command.UserRequest) (*model.User, error)
		DeleteUser(id uint) error

		GetUserMessages(id uint) ([]*model.UserMessage, error)
	}

	userService struct {
		userRepository        repository.UserRepository
		userMessageRepository repository.UserMessageRepository
		logger                *logrus.Logger
	}
)

func (u userService) GetUserMessages(id uint) ([]*model.UserMessage, error) {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErr.NewUserNotFoundError()
		}

		return nil, err
	}

	return u.userMessageRepository.FindByUserID(user.ID)
}

// UpdateUser updates a user
func (u userService) UpdateUser(id uint, com *command.UserRequest) (*model.User, error) {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErr.NewUserNotFoundError()
		}

		return nil, err
	}

	if user.Email != com.Email {
		exists, err := u.userRepository.FindByEmail(com.Email)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		if exists != nil {
			return nil, appErr.NewUserAlreadyExistsError()
		}
	}

	user.Name = com.Name
	user.Email = com.Email
	user.PhoneNumber = com.PhoneNumber

	if err := u.userRepository.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user
func (u userService) DeleteUser(id uint) error {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return appErr.NewUserNotFoundError()
		}

		return err
	}

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
