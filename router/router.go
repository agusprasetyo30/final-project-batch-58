package router

import (
	"final-project/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	api := r.Group("/api")

	// Class / Kelas
	class := api.Group("/class")
	class.GET("/", controllers.GetAllClasses)
	class.POST("/", controllers.InsertClass)
	class.GET("/:id", controllers.GetClass)
	class.PUT("/:id", controllers.UpdateClass)
	class.DELETE("/:id", controllers.DeleteClass)

	// mapel / Mata Pelajaran
	mapel := api.Group("/mata-pelajaran")
	mapel.GET("/", controllers.GetAllMapel)
	mapel.POST("/", controllers.InsertMapel)
	mapel.GET("/:id", controllers.GetMapel)
	mapel.PUT("/:id", controllers.UpdateMapel)
	mapel.DELETE("/:id", controllers.DeleteMapel)

	// class.GET("/:id", controllers.getClass)
	// class.POST("/:id/edit", controllers.InsertClass)
	// class.POST("/:id/aaaaaa", controllers.InsertClass)

	// mapel := api.Group("/mapel")
	// mapel.POST("/", controllers.InsertClass)

	// r.Use(middleware.AuthMiddleware())

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
