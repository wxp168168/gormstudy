package main

import (
	"fmt"
	"gormstudy/common"
	"time"

	"gorm.io/gorm"
)

type UserAfterSave struct {
	ID        uint
	Name      string `gorm:"size:30"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *UserAfterSave) AfterSave(tx *gorm.DB) (err error) {
	fmt.Println("AfterSave hook is triggered")
	// 在这里可以执行一些与保存操作相关的后续逻辑
	return nil
}

/**
以下是一个使用 Gorm 的 AfterSave 钩子的示例，其中结构体名为 UserAfterSave，该钩子在执行保存（插入或更新）操作之后触发
*/

func main() {
	db := common.MysqlCon()

	// 创建表
	db.AutoMigrate(&UserAfterSave{})

	// 创建用户
	user := UserAfterSave{Name: "John Doe"}
	db.Save(&user) // AfterSave 钩子会在这里触发

	// 更新用户
	user.Name = "Jane Doe"
	db.Save(&user) // AfterSave 钩子会在这里触发
}
