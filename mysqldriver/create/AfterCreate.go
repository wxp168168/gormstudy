package main

import (
	"fmt"
	"gormstudy/common"
	"time"

	"gorm.io/gorm"
)

type UserAfterCreate struct {
	ID        uint
	Name      string `gorm:"size:30"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *UserAfterCreate) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterCreate hook is triggered")
	// 在这里可以执行一些与创建操作相关的后续逻辑
	return nil
}

/**
以下是一个使用 Gorm 的 AfterCreate 钩子的示例，其中结构体名为 UserAfterCreate，该钩子在执行创建（插入）操作之后触发
*/

func main() {
	db := common.MysqlCon()

	// 创建表
	db.AutoMigrate(&UserAfterCreate{})

	// 创建用户
	user := UserAfterCreate{Name: "John Doe"}
	db.Create(&user) // AfterCreate 钩子会在这里触发
}
