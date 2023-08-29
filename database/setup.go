package database

import (
	"task-5-vix-btpns/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/dashboard"))
	if err != nil{
		panic(err)
	}

	database.AutoMigrate(&models.Photo{}, &models.User{})

	DB = database
}