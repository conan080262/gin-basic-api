package configs

import (
	"fmt"
	"os"

	"github.com/conan080262/gin-basic-api.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB Connect Lost")
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println("DB Connect Successfully")

	// Migration
	db.AutoMigrate(&models.User{})

	DB = db
}
