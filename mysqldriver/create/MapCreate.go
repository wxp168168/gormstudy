package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserMap struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gormstudy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("无法连接到数据库")
	}

	// 自动迁移模式，将 `UserMap` 结构体映射到数据库表
	db.AutoMigrate(&UserMap{})

	// 创建记录
	db.Model(&UserMap{}).Create(map[string]interface{}{
		"Name": "jinzhu",
		"Age":  18,
	})

	// 查询记录
	var user UserMap
	db.First(&user, "Name = ?", "jinzhu")
	fmt.Println(user)

	// 更新记录
	db.Model(&user).Update("Age", 20)

	// 删除记录
	//db.Delete(&user)
}
