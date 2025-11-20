package configs

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"log"
	"go-proj/models"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/fiberdb?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
        log.Fatal(err)
    }

	DB.AutoMigrate(&models.User{})
}