package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := "root@tcp(127.0.0.1:3306)/api_fiber?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return DB
}
