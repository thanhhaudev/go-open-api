package model

import (
	"gorm.io/gorm"
	"time"
)

type Tenant struct {
	ID        uint `gorm:"primarykey"`
	Scope     string
	Name      string
	AppKey    string
	AppSecret string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Tenant) TableName() string {
	return "tenants"
}
