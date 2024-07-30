package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	dns string
}

func (db *DB) NewDB(dns string) *DB {
	return &DB{
		dns: dns,
	}
}

func (db *DB) MySQL() *gorm.DB {
	open, err := gorm.Open(mysql.Open(db.dns), &gorm.Config{})
	if err != nil {
		panic("Can't connect to database, error=" + err.Error()) // TODO: custom error
	}

	return open
}
