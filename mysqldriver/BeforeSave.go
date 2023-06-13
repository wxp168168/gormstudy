package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserBeforeSave struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

/**
我们创建了一个新的用户并调用 Save 方法保存。此时，BeforeSave 钩子会被触发并打印出相应的信息。
接下来，我们更新了用户的名称，并再次调用 Save 方法，这将再次触发 BeforeSave 钩子

BeforeSave hook is triggered
BeforeSave hook is triggered

BeforeSave|BeforeCreate|AfterSave|AfterCreate这些函数的作用类似
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
