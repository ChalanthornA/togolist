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

	auth := app.Group("/auth")
	auth.Post("/signup", controller.CreateUser)
	auth.Post("/signin", controller.Signin)
	auth.Get("/getProfile", controller.AuthorizationRequired(), controller.GetUserProfile)

	todo := app.Group("/todo", controller.AuthorizationRequired())
	todo.Post("/createNewInv", controller.CreateNewInventory)
	todo.Post("/createNewTodo", controller.CreateTodo)
	todo.Get("/getTodo", controller.GetUserTodo)
	todo.Delete("/deleteTodo/:id", controller.DeleteTodo)
	todo.Patch("/updateTodo", controller.UpdateUserTodoList)
}