package controller

import (
	"log"
	"fmt"
	"github.com/ChalanthornA/katradebutgo/config"
	"github.com/ChalanthornA/katradebutgo/entity"
	"github.com/gofiber/fiber/v2"
)

func CreateNewInventory(c *fiber.Ctx) error{
	id := GetIdFromToken(c)
	queryUser, err := GetUserById(id)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": queryUser,
		})
	}
	inventory := new(entity.Inventory)
	if err := c.BodyParser(inventory); err != nil {
		log.Fatal(err)
		c.Status(400).JSON(&fiber.Map{
            "success": false,
            "message": err,
        })
	}
	fmt.Printf("%+v\n", queryUser);
	newInventory := entity.Inventory{
		Name: inventory.Name,
		Detail: inventory.Detail,
		User: queryUser,
	}
	config.DB.Create(&newInventory);
	c.Status(200).JSON(&fiber.Map{
		"success": true,
	})
	return nil;
}

func GetAllUserInventory(c *fiber.Ctx)error{
	id := GetIdFromToken(c)
	inventories := new([]entity.Inventory)
	if err := config.DB.Where("user_refer = ?", id).Find(&inventories).Error; err == nil{
		c.Status(200).JSON(&fiber.Map{
			"success": true,
			"inventories": inventories,
		})
	}else{
		c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
		})
	}
	return nil
}