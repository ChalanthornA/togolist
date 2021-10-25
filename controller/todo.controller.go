package controller

import (
	"log"
	"time"
	"github.com/ChalanthornA/katradebutgo/config"
	"github.com/ChalanthornA/katradebutgo/entity"
	"github.com/gofiber/fiber/v2"
)


func CreateTodo(c *fiber.Ctx) error{
	id := GetIdFromToken(c)
	queryUser, err := GetUserById(id)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": queryUser,
		})
	}
	bodyTodo := new(entity.Todo)
	if err := c.BodyParser(&bodyTodo); err != nil {
		log.Fatal(err)
		c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
            "success": false,
            "message": err,
        })
	}
	newTodo := entity.Todo{
		Todo: bodyTodo.Todo,
		Create: time.Now(),
		User: queryUser,
	}
	config.DB.Create(&newTodo)
	c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": newTodo,
	})
	return nil
}

func GetUserTodo(c *fiber.Ctx) error{
	id := GetIdFromToken(c)
	queryUser, err := GetUserById(id)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
		})
	}
	queryTodo := new([]entity.Todo)
	if err := config.DB.Where("user_refer = ?", queryUser.ID).Find(&queryTodo).Error; err != nil{
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
		})
	}
	c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"TodoList": queryTodo,
	})
	return nil
}

func DeleteTodo(c *fiber.Ctx) error{
	return nil
}

func GetGuestTodoList(c *fiber.Ctx) error{
	return nil
}

func PostGuestTodoList(c *fiber.Ctx) error{
	return nil
}