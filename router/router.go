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

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())

	// Class / Kelas
	class := api.Group("/class")
	class.GET("/", controllers.GetAllClasses)
	class.POST("/", middleware.RoleMiddleware("ADMIN"), controllers.InsertClass)
	class.GET("/:id", controllers.GetClass)
	class.PUT("/:id", middleware.RoleMiddleware("ADMIN"), controllers.UpdateClass)
	class.DELETE("/:id", middleware.RoleMiddleware("ADMIN"), controllers.DeleteClass)

	// Course / Mata Pelajaran
	course := api.Group("/course")
	course.GET("/", controllers.GetAllCourse)
	course.POST("/", middleware.RoleMiddleware("ADMIN"), controllers.InsertCourse)
	course.GET("/:id", controllers.GetCourse)
	course.PUT("/:id", middleware.RoleMiddleware("ADMIN"), controllers.UpdateCourse)
	course.DELETE("/:id", middleware.RoleMiddleware("ADMIN"), controllers.DeleteCourse)

	// Student / Siswa
	student := api.Group("/student")
	student.GET("/", controllers.GetAllStudents)
	student.GET("/class/:class_id", controllers.GetStudentByClass)
	student.GET("/:id/cources", controllers.GetStudentCourse)
	student.POST("/", middleware.RoleMiddleware("ADMIN"), controllers.InsertStudent)
	student.GET("/:id", controllers.GetStudent)
	student.PUT("/:id", middleware.RoleMiddleware("ADMIN"), controllers.UpdateStudent)
	student.DELETE("/:id", middleware.RoleMiddleware("ADMIN"), controllers.DeleteStudent)

	// Assessment / Penilaian
	assessment := api.Group("/assessment")
	assessment.GET("/", controllers.GetAllAssessment)
	assessment.POST("/", middleware.RoleMiddleware("ADMIN"), controllers.InsertAssessment)
	assessment.GET("/:id", controllers.GetAssessment)
	assessment.PUT("/:id", middleware.RoleMiddleware("ADMIN"), controllers.UpdateAssessment)
	assessment.DELETE("/:id", middleware.RoleMiddleware("ADMIN"), controllers.DeleteAssessment)

	return r
}
