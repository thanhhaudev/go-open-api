package model

import (
	"gorm.io/gorm"

	"time"
)

type Message struct {
	ID        uint           `gorm:"primaryKey" json:"id" example:"1"`
	Subject   string         `json:"subject" example:"Hello"`
	Content   string         `json:"content" example:"Hello, how are you?"`
	CreatedAt time.Time      `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2021-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggerignore:"true"`

	Sender   *User   `gorm:"foreignKey:ID" json:"sender"` // read-only
	Receiver []*User `gorm:"many2many:user_messages;" json:"receiver"`
}
