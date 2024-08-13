package mysql

import (
	"github.com/thanhhaudev/openapi-go/app/model"
	"github.com/thanhhaudev/openapi-go/app/repository"
	"gorm.io/gorm"
)

type userMessageRepository struct {
	gorm *gorm.DB
}

// FindByUserID finds user messages by user ID
func (u userMessageRepository) FindByUserID(userId uint) ([]*model.UserMessage, error) {
	var userMessages []*model.UserMessage

	err := u.gorm.
		Select("user_messages.*, messages.subject, messages.content").
		Joins("JOIN messages ON user_messages.message_id = messages.id").
		Where("user_id = ?", userId).
		Find(&userMessages).Error
	if err != nil {
		return nil, err
	}

	return userMessages, nil
}

// FindByID finds a user message by its ID
func (u userMessageRepository) FindByID(userId, id uint) (*model.UserMessage, error) {
	userMessage := &model.UserMessage{}

	err := u.gorm.
		Select("user_messages.*, messages.subject, messages.content").
		Joins("JOIN messages ON user_messages.message_id = messages.id").
		Where("user_messages.user_id = ? AND messages.id = ?", userId, id).First(userMessage).Error
	if err != nil {
		return nil, err
	}

	return userMessage, nil
}

func (u userMessageRepository) Create(userMessage *model.UserMessage) error {
	err := u.gorm.Create(userMessage).Error
	if err != nil {
		return err
	}

	return nil
}

func (u userMessageRepository) Update(userMessage *model.UserMessage) error {
	err := u.gorm.Save(userMessage).Error
	if err != nil {
		return err
	}

	return nil
}

func (u userMessageRepository) Delete(userId, id uint) error {
	err := u.gorm.Where("user_id = ? AND id = ?", userId, id).Delete(&model.UserMessage{}).Error
	if err != nil {
		return err
	}

	return nil
}

// NewUserMessageRepository creates a new user message repository
func NewUserMessageRepository(gorm *gorm.DB) repository.UserMessageRepository {
	return &userMessageRepository{
		gorm: gorm,
	}
}
