package controllers

import (
	"fmt"
	"log"
	"github.com/ChalanthornA/katradebutgo/config"
	// "github.com/ChalanthornA/katradebutgo/database"
	"github.com/ChalanthornA/katradebutgo/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err;
}

func CompareHashPassword(hash string, password string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil;
}

func CreateUser(c *fiber.Ctx) error{
	user := new(entity.User);
	queryUser := new(entity.User);
	if err := c.BodyParser(user); err != nil {
		log.Fatal(err)
		c.Status(400).JSON(&fiber.Map{
            "success": false,
            "message": err,
        })
	}
	if err := config.DB.Where("email = ?", user.Email).First(&queryUser).Error; err == nil{
		fmt.Printf("%+v\n", queryUser)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "This email is already used",
		})
	} 
	hashPassword, _ := HashPassword(user.Password)
	newUser := entity.User{
		ID : uuid.NewV4(), 
		Name: user.Name,
		Email: user.Email,
		Password: hashPassword,
	}
	config.DB.Create(&newUser)

	// res, err := database.DB.Query("INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)", 1, user.Name, user.Email, user.Password)
	// if err != nil {
    //     c.Status(500).JSON(&fiber.Map{
    //         "success": false,
    //         "message": err,
    //     })
    //     return nil
    // }
	// log.Println(res)
	c.Status(200).JSON(&fiber.Map{
		"success": true,
	})
	return nil;
}

func Signin(c *fiber.Ctx) error{
	loginForm := new(entity.LoginForm);
	user := new(entity.User);
	if err := c.BodyParser(loginForm); err != nil {
		log.Fatal(err)
		c.Status(400).JSON(&fiber.Map{
            "success": false,
            "message": err,
        })
	}
	if err := config.DB.Where("email = ?", loginForm.Email).First(&user).Error; err != nil{
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "Cannot find user",
		})
	}
	result := CompareHashPassword(user.Password, loginForm.Password);
	if result{
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
		})
	}
	return c.Status(400).JSON(&fiber.Map{
		"success": false,
		"message": "Wrong password",
	})
}