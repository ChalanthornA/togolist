package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ChalanthornA/katradebutgo/controllers"
)

func sayhi(c *fiber.Ctx) error{
	return c.SendString("HiHi")
}

func SetUpRoute(app *fiber.App){
	app.Get("/", sayhi)
	app.Post("/signup", controllers.CreateUser)
	app.Post("/signin", controllers.Signin)
	app.Get("/getProfile",controllers.AuthorizationRequired(), controllers.GetUserProfile)
}