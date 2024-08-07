package model

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID        int            `gorm:"primaryKey" json:"-" swaggerignore:"true"`
	Subject   string         `json:"subject" example:"Hello"`
	Content   string         `json:"content" example:"Hello, how are you?"`
	Sender    *User          `gorm:"foreignKey:ID" json:"sender"`
	CreatedAt time.Time      `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updatedAt" example:"2021-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
}
