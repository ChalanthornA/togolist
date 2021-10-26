package config

import (
	"github.com/ChalanthornA/togolist/entity"
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
		log.Fatal(errEnv)
		log.Fatalf("Error loading .env file")
	}
	var err error;
	dsn := os.Getenv("DSN")
	fmt.Printf("this is %s\n", dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Fail to connect")
		os.Exit(100);
	}
	fmt.Println("connect to db")
	DB.AutoMigrate(&entity.User{}, &entity.Todo{})
	return DB;
}