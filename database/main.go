package database

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Init(config ...Config) (string, error) {
	var err error
	var db *gorm.DB
	for _, item := range config {
		switch item.Driver {
		case "mysql":
			db, err = gorm.Open(mysql.Open(item.Dsn), item.Config)
		case "sqlsrv":
			db, err = gorm.Open(sqlserver.Open(item.Dsn), item.Config)
		}
		if err != nil {
			return item.Name, err
		}
		Session[item.Name] = session{Name: item.Name, DB: db}
	}
	return "", nil
}
func InitRedis(cfg RedisConfig) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// Test connection
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}
