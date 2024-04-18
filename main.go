package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique"`
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {

	db, err := gorm.Open(mysql.Open("ชื่อผู้ใช้:รหัสผ่าน(127.0.0.1:3306)/ชื่อฐานข้อมูล?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&User{})

	user := User{
		Name:  "Sirasit",
		Email: "sirasit@example.com",
		Age:   25,
	}
	result := db.Create(&user)
	if result.Error != nil {
		panic("failed to create user")
	}
}
