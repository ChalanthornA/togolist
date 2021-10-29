package controller

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ChalanthornA/togolist/config"
	"github.com/ChalanthornA/togolist/entity"
	"github.com/ChalanthornA/togolist/types"
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
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
            "success": false,
            "message": err,
        })
	}
	newTodo := entity.Todo{
		Todo: bodyTodo.Todo,
		Description: bodyTodo.Description,
		Create: time.Now(),
		UserRefer: queryUser.ID,
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
		"Name": queryUser.Name,
		"TodoList": queryTodo,
	})
	return nil
}

func DeleteTodo(c *fiber.Ctx) error{
	userId := int64(GetIdFromToken(c))
	deleteId, err := strconv.Atoi(c.Params("id"))
	if err != nil{
		log.Fatal(err)
	}
	queryTodo := new(entity.Todo)
	if err := config.DB.First(&queryTodo, deleteId).Error; err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "Can't find this todo.",
		})
	}
	if queryTodo.UserRefer == userId {
		if err := config.DB.Delete(&entity.Todo{}, deleteId).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
			})
		}
	}else{
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "This is not your todo.",
		})
	}
	c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
	return nil
}

func UpdateUserTodoList(c *fiber.Ctx) error{
	userId := int64(GetIdFromToken(c))
	updateMessage := new(types.UpdateMessage)
	if err := c.BodyParser(&updateMessage); err != nil{
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
            "success": false,
            "message": err,
        })
	}
	queryTodo := new(entity.Todo)
	if err := config.DB.First(&queryTodo, updateMessage.ID).Error; err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "Can't find this todo.",
		})
	}
	if userId == queryTodo.UserRefer {
		fmt.Printf("%s\n", updateMessage.Description)
		if updateMessage.Todo != "" && updateMessage.Description != ""{
			if err := config.DB.Model(&entity.Todo{}).Where("id = ?", updateMessage.ID).Updates(entity.Todo{Todo: updateMessage.Todo, Description: updateMessage.Description}).Error; err != nil{
				return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
					"success": false,
				})
			}
		}else if updateMessage.Todo != "" {
			if err := config.DB.Model(&entity.Todo{}).Where("id = ?", updateMessage.ID).Update("todo", updateMessage.Todo).Error; err != nil{
				return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
					"success": false,
				})
			}
		}else {
			if err := config.DB.Model(&entity.Todo{}).Where("id = ?", updateMessage.ID).Update("description", updateMessage.Description).Error; err != nil{
				return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
					"success": false,
				})
			}
		}
	}else{
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "This is not your todo.",
		})
	}
	c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
	return nil
}

func GetGuestTodoList(c *fiber.Ctx) error{
	return nil
}

func PostGuestTodoList(c *fiber.Ctx) error{
	return nil
}