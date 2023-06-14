package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlCon() (db *gorm.DB) {
	dsn := "root:root@tcp(127.0.0.1:3306)/gormstudy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}
