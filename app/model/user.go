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
	UpdatedAt   time.Time      `json:"updated_at" example:"2021-01-01T00:00:00Z"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
}

func (u User) TableName() string {
	return "users"
}

type UserMessage struct {
	UserID    uint           `gorm:"primarykey" json:"userId" example:"1"`
	MessageID uint           `gorm:"primarykey" json:"messageId" example:"1"`
	Subject   string         `gorm:"embedded;embeddedPrefix:message_" json:"subject" example:"Hello"`
	Content   string         `gorm:"embedded;embeddedPrefix:message_" json:"content" example:"Hello, how are you?"`
	IsRead    bool           `json:"is_read" example:"false"`
	ReadAt    *time.Time     `json:"readAt" example:"2021-01-01T00:00:00Z"`
	CreatedAt time.Time      `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2021-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
}

func (u UserMessage) TableName() string {
	return "user_messages"
}
