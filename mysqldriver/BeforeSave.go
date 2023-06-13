package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserBeforeSave struct {
	ID        uint
	Name      string `gorm:"size:30"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

/**
以下是一个使用 Gorm 的 BeforeSave 钩子的示例，其中结构体名为 UserBeforeSave，该钩子在执行保存（插入或更新）操作之前触发
*/

func (u *UserBeforeSave) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Println("BeforeSave hook is triggered")
	// 在这里可以执行一些预处理或验证逻辑
	return nil
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/gormstudy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 创建表
	db.AutoMigrate(&UserBeforeSave{})

	// 创建用户
	user := UserBeforeSave{Name: "John Doe"}
	db.Save(&user) // BeforeSave 钩子会在这里触发

	// 更新用户
	user.Name = "Jane Doe"
	db.Save(&user) // BeforeSave 钩子会在这里触发
}
