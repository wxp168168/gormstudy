package main

//######## 用指定的字段创建记录 Select Create ########

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PersonSelect struct {
	ID        uint
	FirstName string `gorm:"size:30"`
	LastName  string `gorm:"size:30"`
	Age       int    `gorm:"type:tinyint(3)"`
}

func main() {
	// 建立数据库连接
	dsn := "root:root@tcp(localhost:3306)/gormstudy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// 错误处理
		fmt.Println("数据库连接错误:", err)
		return
	}

	// 迁移表格（如果尚未创建）
	err = db.AutoMigrate(&PersonSelect{})
	if err != nil {
		// 错误处理
		fmt.Println("数据库迁移错误:", err)
		return
	}

	// 创建personSelect记录
	personSelect := &PersonSelect{
		FirstName: "John",
		LastName:  "Doe",
		Age:       101,
	}

	// 选择指定字段并创建记录
	err = db.Select("first_name", "Age").Create(&personSelect).Error
	if err != nil {
		fmt.Println("创建记录错误:", err)
		return
	}

	fmt.Println("创建记录成功!")
}
