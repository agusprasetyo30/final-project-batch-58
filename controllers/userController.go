package controllers

import (
	"final-project/database"
	"final-project/helper"
	"final-project/model"
	"final-project/repository"
	"final-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	// authService services.AuthService
	authRepository repository.UserRepository
}

// func NewAuthController(authService services.AuthService) *AuthController {
func NewAuthController(ar repository.UserRepository) *AuthController {
	// return &AuthController{authService: authService}
	return &AuthController{authRepository: ar}
}

func Login(ctx *gin.Context) {
	userRepository := repository.NewUserRepository(database.DbConnection)
	authService := services.NewAuthService(userRepository)

	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username dan Password tidak boleh kosong"})
		return
	}

	user, err := authService.Authenticate(username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing credentials"})
		return
	}

	if user == nil || !helper.VerifyPassword(password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username or Password"})
		return
	}

	// Login successful, return user data or token
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func Register(c *gin.Context) {
	user := &model.User{}
	err := c.BindJSON(user)

	if err != nil {
		panic(err)
	}

	user.Password = helper.GeneratePassword(user.Password)
	data, err := repository.CreateUser(database.DbConnection, *user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, data)
}
