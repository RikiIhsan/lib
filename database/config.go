package database

import "gorm.io/gorm"

type Config struct {
	Name    string
	Dsn     string
	Driver  string
	Config *gorm.Config
}

type session struct {
	Name string
	DB   *gorm.DB
}

var Session = make(map[string]session)
