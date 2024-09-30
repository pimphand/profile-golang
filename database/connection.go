package database

import (
	"gin-gonic-gorm/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	if config.DB_Driver == "mysql" {
		DB, err = gorm.Open(mysql.Open(config.DB_Username+":"+config.DB_Password+"@tcp("+config.DB_Host+":"+config.DB_Port+")/"+config.DB_Name+"?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	}

	if err != nil {
		panic("Failed to connect to database!")
	}
}
