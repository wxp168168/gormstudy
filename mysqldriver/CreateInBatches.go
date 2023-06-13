package main

//######## 批量插入 CreateInBatches ########

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string `gorm:"size:30"`
	Birthday  string `gorm:"size:10"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/gormstudy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto Migrate the User model
	db.AutoMigrate(&User{})

	// 创建一些用户数据
	users := []User{
		{Name: "Alice", Birthday: time.Now().Format("2006-01-02")},
		{Name: "Bob", Birthday: time.Now().AddDate(0, -1, 0).Format("2006-01-02")},
		{Name: "Charlie", Birthday: time.Now().AddDate(0, 0, -7).Format("2006-01-02")},
		// 其他用户...
	}

	// 使用CreateInBatches批量创建用户记录
	result := db.CreateInBatches(&users, 2) // 每批次创建3个记录

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("创建的用户数量:", result.RowsAffected)
}
