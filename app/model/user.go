package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Email       string         `json:"email"`
	Name        string         `json:"name"`
	PhoneNumber *string        `gorm:"string" json:"phone_number"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
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
