package configs

import (
	"fmt"
	"go-jwt/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf(
		"host=%s user=postgres password=%s dbname=latihan port=5434 sslmode=disable TimeZone=Asia/Shanghai",
		HOST,
		PASS,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})
	DB = db
	log.Println("Database connected successfully")
}
