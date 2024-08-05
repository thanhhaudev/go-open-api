package model

import (
	"gorm.io/gorm"
	"time"
)

type Tenant struct {
	ID        uint `gorm:"primarykey"`
	Scope     string
	Name      string
	ApiKey    string
	ApiSecret string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Tenant) TableName() string {
	return "tenants"
}
