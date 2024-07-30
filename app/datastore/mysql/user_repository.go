package mysql

import (
	"github.com/thanhhaudev/openapi-go/app/model"
	"github.com/thanhhaudev/openapi-go/app/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	gorm *gorm.DB
}

func (u userRepository) FindAll() ([]*model.User, error) {
	var r []*model.User

	err := u.gorm.Find(&r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (u userRepository) FindByID(id int64) (*model.User, error) {
	user := &model.User{}

	err := u.gorm.First(user, id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userRepository) Create(user *model.User) error {
	err := u.gorm.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u userRepository) Update(user *model.User) error {
	err := u.gorm.Save(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u userRepository) Delete(id int64) error {
	err := u.gorm.Delete(&model.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

// NewUserRepository creates a new user repository
func NewUserRepository(gorm *gorm.DB) repository.UserRepository {
	return userRepository{
		gorm: gorm,
	}
}
