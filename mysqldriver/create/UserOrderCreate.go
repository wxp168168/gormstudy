package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserCre struct {
	ID     uint
	Name   string     `gorm:"size:30"`
	Orders []OrderCre `gorm:"foreignkey:UserID"` // 定义外键关系
}

type OrderCre struct {
	ID     uint
	UserID uint
	Amount float64
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gormstudy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("无法连接到数据库")
	}

	// 在数据库中创建表
	db.AutoMigrate(&UserCre{})
	db.AutoMigrate(&OrderCre{})

	user := UserCre{
		Name: "John Doe",
		Orders: []OrderCre{
			{Amount: 10.5},
			{Amount: 15.2},
		},
	}

	// 创建用户和关联的订单
	db.Create(&user)
}
