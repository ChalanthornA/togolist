package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ChalanthornA/katradebutgo/controller"
)

// SayHi godoc
// @Summary SayHi
// @Description LetKatrade say hi
// @Success 200 
// @Router / [get]
func sayhi(c *fiber.Ctx) error{
	return c.Status(200).JSON(&fiber.Map{
		"message" : "hihi",
	})
}

func SetUpRoute(app *fiber.App){
	app.Get("/", sayhi)
	app.Post("/signup", controller.CreateUser)
	app.Post("/signin", controller.Signin)
	app.Use(controller.AuthorizationRequired())
	app.Get("/getProfile", controller.GetUserProfile)
	app.Post("/createNewInv", controller.CreateNewInventory)
	app.Get("/getUserInv", controller.GetAllUserInventory)
	app.Post("/createNewTodo", controller.CreateTodo)
	app.Get("/getTodo", controller.GetUserTodo)
}