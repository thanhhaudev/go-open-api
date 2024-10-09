package mysql

import (
	"github.com/thanhhaudev/go-open-api/app/model"
	"github.com/thanhhaudev/go-open-api/app/repository"
	"gorm.io/gorm"
)

type messageRepository struct {
	gorm *gorm.DB
}

// FindByID finds a message by its ID
func (m messageRepository) FindByID(id uint) (*model.Message, error) {
	message := &model.Message{}

	err := m.gorm.
		Preload("Sender").
		Preload("Receivers").
		First(message, id).Error
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (m messageRepository) Create(message *model.Message) error {
	err := m.gorm.Omit("Sender").Create(message).Error
	if err != nil {
		return err
	}

	return nil
}

func (m messageRepository) Update(message *model.Message) error {
	err := m.gorm.Save(message).Error
	if err != nil {
		return err
	}

	return nil
}

func (m messageRepository) Delete(id uint) error {
	err := m.gorm.Delete(&model.Message{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func NewMessageRepository(gorm *gorm.DB) repository.MessageRepository {
	return &messageRepository{
		gorm: gorm,
	}
}
