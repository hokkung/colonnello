package entity

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ProvideDB() (*gorm.DB, func(), error) {
	dsn := "root:@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, func() {}, nil
}
