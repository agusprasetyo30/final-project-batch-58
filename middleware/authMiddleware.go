package middleware

import (
	"final-project/database"
	"final-project/helper"
	"final-project/model"
	"final-project/repository"
	"final-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRepository := repository.NewUserRepository(database.DbConnection)
		authService := services.NewAuthService(userRepository)

		username, password, ok := ctx.Request.BasicAuth()

		if username == "" || password == "" || !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing credentials"})
			ctx.Abort()
			return
		}

		// Call the auth controller to authenticate
		user, err := authService.Authenticate(username)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing credentials"})
			ctx.Abort()
			return
		}

		if user == nil || !helper.VerifyPassword(password, user.Password) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username or Password"})
			return
		}

		// Set the authenticated user in the context
		ctx.Set("user", user)
		ctx.Next()
	}
}

func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			helper.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		if user.(*model.User).Role != role {
			helper.RespondWithError(c, http.StatusForbidden, "Forbidden Role")
			c.Abort()
			return
		}

		c.Next()
	}
}
