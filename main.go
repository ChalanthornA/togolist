package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ChalanthornA/togolist/router"
	"github.com/ChalanthornA/togolist/config"
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/ChalanthornA/togolist/docs"
	"os"
)


// @title Katrade Api
// @version 1.0
// @description This is a sample swagger for Katrade
// @contact.name Katrade Backend
// @contact.email youremail@provider.com
// @host localhost:8080
// @BasePath /
func main(){
	// if err := database.Connect(); err != nil{
	// 	log.Fatal(err)
	// }
	config.Connect();

	app := fiber.New();

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

	router.SetUpRoute(app);

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app.Listen(":" + port)
}