package router

import (
	"final-project/controllers"
	"final-project/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	r.Use(middleware.AuthMiddleware())

	// api := r.Group("/api")
	// api.Use(middleware.AuthMiddleware())
	// {
	// 	api.GET("/protected", protectedHandler)
	// 	api.GET("/admin", middleware.RoleMiddleware("admin"), adminHandler)
	// }

	return r
}

// func protectedHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"message": "This is a protected route"})
// }

// func adminHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"message": "This is an admin route"})
// }
