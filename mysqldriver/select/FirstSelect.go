package main

import (
	"fmt"
	"gormstudy/common"
)

type UserFirst struct {
	ID   uint
	Name string
	Age  uint
}

func main() {
	db := common.MysqlCon()

	// 自动迁移模式，创建users表
	db.AutoMigrate(&UserFirst{})

	uf := []UserFirst{{Name: "tom", Age: 18}, {Name: "mactom", Age: 31}}
	db.Create(&uf)

	// 获取第一条记录（主键升序）
	// SELECT * FROM user_firsts ORDER BY id LIMIT 1;
	var user UserFirst
	if err := db.First(&user).Error; err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("First user:", user)
	}
}
