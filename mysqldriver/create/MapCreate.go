package main

import (
	"fmt"
	"gormstudy/common"
)

type UserMap struct {
	ID    int
	Name  string `gorm:"size:30"`
	Email string `gorm:"size:60"`
	Age   int    `gorm:"type:tinyint(3)"`
}

func main() {
	db := common.MysqlCon()

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
