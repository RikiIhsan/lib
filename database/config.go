package database

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Config struct {
	Name   string
	Dsn    string
	Driver string
	Config *gorm.Config
}
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type session struct {
	Name string
	DB   *gorm.DB
}

var (
	Session       = make(map[string]session)
	RedisClient   *redis.Client
)
