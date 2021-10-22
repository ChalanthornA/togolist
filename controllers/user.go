package controllers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ChalanthornA/katradebutgo/config"
	"github.com/ChalanthornA/katradebutgo/entity"
	"github.com/ChalanthornA/katradebutgo/types"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
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
	loginForm := new(types.LoginForm);
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
		accessToken, err := CreateToken(user.ID, user.Name)
		if err != nil{
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"token": accessToken,
			})
		}
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"token": accessToken,
		})
	}
	return c.Status(400).JSON(&fiber.Map{
		"success": false,
		"message": "Wrong password",
	})
}

func CreateToken(uid uuid.UUID, name string) (string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = uid
    claims["name"] = name
    claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	secret := os.Getenv("SECRET");
	accessToken, err := token.SignedString([]byte(secret))
	return accessToken, err;
}

func AuthError(c *fiber.Ctx, e error) error {
    c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
        "error": "Unauthorized",
        "msg":   e.Error(),
    })
    return nil
}

func AuthSuccess(c *fiber.Ctx) error {
    c.Next()
    return nil
}

func AuthorizationRequired() fiber.Handler {
	secret := os.Getenv("SECRET");
    return jwtware.New(jwtware.Config{
        // Filter:         nil,
        SuccessHandler: AuthSuccess,
        ErrorHandler:   AuthError,
        SigningKey:     []byte(secret),
        // SigningKeys:   nil,
        SigningMethod: "HS256",
        // ContextKey:    nil,
        // Claims:        nil,
        // TokenLookup:   nil,
        // AuthScheme:    nil,
    })
}

func GetUserProfile(c *fiber.Ctx) error{
	user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    sub := claims["sub"].(string)
	queryUser := new(entity.User)
	if err := config.DB.Where("id = ?", sub).First(&queryUser).Error; err == nil{
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": queryUser,
		})
	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{
        "sub": sub,
    })
	return nil;
}