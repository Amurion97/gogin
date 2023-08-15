package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/Data"
	"go-gin/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gogin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Models.User{})

	user := Models.User{Name: "Giang", Age: 15}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(user.ID)

	var firstUser Models.User
	db.First(&firstUser)
	fmt.Println(firstUser)

	var foundUser Models.User
	db.Where(
		"name = ?",
		"Giang",
	).Where("age > ?", 0).First(&foundUser)
	fmt.Println("foundUser:", foundUser)
	//----------------------------------
	router := gin.Default()
	router.GET("/albums", Data.GetAlbums)
	router.GET("/albums/:id", Data.GetAlbumByID)
	router.POST("/albums", Data.PostAlbums)

	router.Run("localhost:8080")
}

//--------------------------------------------------------------------------------
