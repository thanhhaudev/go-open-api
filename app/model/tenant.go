package model

import (
	"gorm.io/gorm"
	"strings"
	"time"
)

type Tenant struct {
	ID        uint   `gorm:"primarykey"`
	Scopes    string `gorm:"column:scopes"`
	Name      string
	ApiKey    string
	ApiSecret string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (t Tenant) GetScopes() []string {
	return strings.Split(t.Scopes, ",")
}

func (t Tenant) TableName() string {
	return "tenants"
}
