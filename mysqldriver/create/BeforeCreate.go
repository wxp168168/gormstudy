package main

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserBeforeCreate struct {
	ID        uint
	Name      string `gorm:"size:30"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *UserBeforeCreate) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate hook is triggered")
	// 在这里可以执行一些预处理或验证逻辑
	return nil
}

/**
以下是一个使用 Gorm 的 BeforeCreate 钩子的示例，其中结构体名为 UserBeforeCreate，该钩子在执行创建（插入）操作之前触发
*/

func main() {
	db := common.MysqlCon()

	// 创建表
	db.AutoMigrate(&UserBeforeCreate{})

	// 创建用户
	user := UserBeforeCreate{Name: "John Doe"}
	db.Create(&user) // BeforeCreate 钩子会在这里触发
}
