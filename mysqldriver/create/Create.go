package main

//######## 批量插入 Create ########

import (
	"fmt"
	"gormstudy/common"
)

type Product struct {
	ID    uint
	Name  string `gorm:"size:30"`
	Price float64
}

func main() {
	// 建立数据库连接
	db := common.MysqlCon()

	// 迁移表格（如果尚未创建）
	err := db.AutoMigrate(&Product{})
	if err != nil {
		// 错误处理
		fmt.Println("数据库迁移错误:", err)
		return
	}

	// 创建包含product对象的切片
	products := []Product{
		{Name: "Product 1", Price: 10.99},
		{Name: "Product 2", Price: 19.99},
		{Name: "Product 3", Price: 14.99},
	}

	// 批量插入product对象
	err = db.Create(&products).Error
	if err != nil {
		// 错误处理
		fmt.Println("批量插入错误:", err)
		return
	}

	fmt.Println("批量插入成功!")
}
