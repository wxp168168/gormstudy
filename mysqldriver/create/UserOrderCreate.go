package main

import (
	"gormstudy/common"
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
	db := common.MysqlCon()

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
	db.Debug().Create(&user)
}
