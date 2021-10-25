package controller

import (
	"github.com/gin-gonic/gin"
)

type AuthController interface{
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct{

}

func NewAuthController() AuthController{
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message": "login",
	})
}

func (c *authController) Register(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message": "register",
	})
}