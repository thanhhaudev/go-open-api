package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primarykey" json:"id" example:"1"`
	Email       string         `json:"email" example:"test@gmail.com"`
	Name        string         `json:"name" example:"test"`
	PhoneNumber *string        `gorm:"string" json:"phone_number" example:"0123456789"`
	CreatedAt   time.Time      `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt   time.Time      `json:"updatedAt" example:"2021-01-01T00:00:00Z"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt" swaggerignore:"true"`
}

func (u User) TableName() string {
	return "users"
}

type UserMessage struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	UserID    uint           `gorm:"primarykey" json:"userId"`
	Message   *Message       `json:"message"`
	Read      bool           `json:"read"`
	ReadAt    *time.Time     `json:"readAt"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (u UserMessage) TableName() string {
	return "user_messages"
}
