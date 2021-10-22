package config

import (
	"github.com/ChalanthornA/katradebutgo/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	// "github.com/go-pg/pg/v10/orm"
	"fmt"
	"os"
	// "github.com/ChalanthornA/katradebutgo/controllers"
	"github.com/joho/godotenv"
	// "gorm.io/gorm"
	// "gorm.io/driver/mysql"
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

// func SetupDBConnection() *gorm.DB{
// 	err := godotenv.Load();
// 	if err != nil{
// 		panic("Fail to load env");
// 	}
// 	dbUser := os.Getenv("DB_USER");
// 	dbPass := os.Getenv("DB_PASS");
// 	dbHost := os.Getenv("DB_HOST");
// 	dbName := os.Getenv("DB_NAME");

// 	var dns string = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName);
// 	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{});
// 	if err != nil{
// 		panic("Fail to connect to DB");
// 	}
// 	// db.AutoMigrate();
// 	return db;
// }

// func CloseDBConnection(db *gorm.DB){
// 	dbSQL, err := db.DB()
// 	if err != nil {
// 		panic("Fail to close DB");
// 	}
// 	dbSQL.Close();
// }