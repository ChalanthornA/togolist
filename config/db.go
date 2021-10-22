package config

import (
	"github.com/ChalanthornA/katradebutgo/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

var DB *gorm.DB;

func Connect() *gorm.DB{
	errEnv := godotenv.Load();
	if errEnv != nil {
		log.Fatalf("Error loading .env file")
	}
	user := os.Getenv("USER");
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD");
	port := os.Getenv("PORT");
	var err error;
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", user, password, dbname, port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Fail to connect")
		os.Exit(100);
	}
	fmt.Println("connect to db")
	DB.AutoMigrate(&entity.User{}, &entity.Inventory{})
	return DB;
}