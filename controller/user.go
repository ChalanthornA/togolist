package controller

import (
	"fmt"
	"log"
	"os"
	"time"
	// "strconv"
	"github.com/ChalanthornA/katradebutgo/config"
	"github.com/ChalanthornA/katradebutgo/entity"
	"github.com/ChalanthornA/katradebutgo/types"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
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
	if err := config.DB.Where("username = ?", user.Username).First(&queryUser).Error; err == nil{
		fmt.Printf("%+v\n", queryUser)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "This username is already used",
		})
	} 
	hashPassword, _ := HashPassword(user.Password)
	newUser := entity.User{
		Firstname: user.Firstname,
		Lastname: user.Lastname,
		Username: user.Username,
		Password: hashPassword,
	}
	config.DB.Create(&newUser)
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
	fmt.Printf("%v\n", loginForm);
	if err := config.DB.Where("username = ?", loginForm.Username).First(&user).Error; err != nil{
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "Cannot find user",
		})
	}
	result := CompareHashPassword(user.Password, loginForm.Password);
	if result{
		accessToken, err := CreateToken(user.ID, user.Firstname)
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

func CreateToken(uid int64, name string) (string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = uid
    claims["name"] = name
    claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
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

func GetIdFromToken(c *fiber.Ctx) float64{
	user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    id := claims["sub"].(float64)
	return id;
}

func GetUserProfile(c *fiber.Ctx) error{
	id := GetIdFromToken(c)
	queryUser, err := GetUserById(id);
	if err == nil{
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": queryUser,
		})
	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": false,
        "sub": id,
    })
	return nil;
}

func GetUserById(id float64) (entity.User, error){
	uid := int(id)
	queryUser := new(entity.User)
	var err error;
	if err = config.DB.Where("id = ?", uid).First(&queryUser).Error; err == nil {
		return *queryUser, err
	}
	return *queryUser, err
}