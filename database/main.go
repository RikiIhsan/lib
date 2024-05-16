package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Init(config ...Config) (error, string) {
	var err error
	var db *gorm.DB
	for _, item := range config {
		switch item.Driver {
		case "mysql":
			db, err = gorm.Open(mysql.Open(item.Dsn), item.Confgis)
		case "sqlsrv":
			db, err = gorm.Open(sqlserver.Open(item.Dsn), item.Confgis)
		}
		if err != nil {
			return err, item.Name
		}
		Session[item.Name] = session{Name: item.Name, DB: db}
	}
	return nil, ""
}
