package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	Conn *gorm.DB
}

func NewDatabase() *Database {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	) // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can't connect to database, error=" + err.Error()) // TODO: custom error
	}

	return &Database{
		Conn: open,
	}
}
