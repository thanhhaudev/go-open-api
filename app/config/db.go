package config

import (
	"fmt"
	"gorm.io/gorm/logger"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Can't connect to database, error=" + err.Error()) // TODO: custom error
	}

	return &Database{
		Conn: open,
	}
}

type RedisStore struct {
	Client *redis.Client
}

// NewRedisStore creates a new RedisStore
func NewRedisStore() *RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_HOST"),
		Password:     os.Getenv("REDIS_PASSWORD"),
		DB:           0,
		DialTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 3,
		ReadTimeout:  time.Second * 3,
		MaxRetries:   3,
	})

	return &RedisStore{
		Client: client,
	}
}
