package main

//######## 创建记录并更新未给出的字段 Omit Create ########

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PersonOmit struct {
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
	err = db.AutoMigrate(&PersonOmit{})
	if err != nil {
		// 错误处理
		fmt.Println("数据库迁移错误:", err)
		return
	}

	// 创建personOmit记录
	personOmit := &PersonOmit{
		FirstName: "John",
		LastName:  "Doe",
		Age:       101,
	}

	// 创建记录并更新未给出的字段
	err = db.Omit("first_name").Create(&personOmit).Error
	if err != nil {
		// 错误处理
		fmt.Println("创建记录错误:", err)
		return
	}

	fmt.Println("创建记录成功!")
}
