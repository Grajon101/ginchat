package main

import (
	"fmt"
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(193.122.105.124:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	user := &models.UserBasic{}
	//utils.InitMySQL()

	//user = models.GetUserByNameAndPwd("abc", "123")

	user.Name = "ball"
	user.PassWord = "123456"
	user.Phone = "1888888888"
	// Create
	db.Create(user)
	//models.TabelName()
	// Read
	//var user UserBasic
	db.First(user, 1) // 根据整型主键查找
	//db.First(&user, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	fmt.Println(db.First(user, 1))
	fmt.Println(user.ID)
	// Update - 将 user 的 price 更新为 200
	db.Model(user).Update("PassWord", "0987654321")
	// Update - 更新多个字段
	//db.Model(&user).Updates(user{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 user
	//db.Delete(&user, 1)
}
