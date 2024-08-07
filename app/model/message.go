package model

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID        int    `json:"id"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
	Sender    *User  `json:"sender"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
