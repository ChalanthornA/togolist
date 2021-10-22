package main

import (
	"github.com/gofiber/fiber/v2"
	// "log"
	// "github.com/ChalanthornA/katradebutgo/database"
	"github.com/ChalanthornA/katradebutgo/router"
	"github.com/ChalanthornA/katradebutgo/config"
)

func main(){
	// if err := database.Connect(); err != nil{
	// 	log.Fatal(err)
	// }

	config.Connect();

	app := fiber.New();

	router.SetUpRoute(app);

	app.Listen(":8080")
}