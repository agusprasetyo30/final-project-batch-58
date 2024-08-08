package main

import (
	"database/sql"
	"final-project/database"
	"final-project/router"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		panic("error loading env file")
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	database.DBMigrate(DB)

	// router := gin.Default()
	// router.Use(middleware.AuthMiddleware())
	// router.POST("/api/login", controllers.Login)

	// config.LoadConfig()

	// router := gin.Default()
	// router.Use(middleware.AuthMiddleware())
	// router.POST("/api/login", controllers.Login)

	// cek := helper.HashPassword("admin")

	// fmt.Println(helper.VerifyPassword("admin", cek))

	router := router.SetupRouter()
	router.Run(":8080")

	// err := bcrypt.CompareHashAndPassword([]byte("$2a$14$8cbHhkS0r2zwJ4br75ILpO/FqVPI.H5d1FRpwee90JhZhtVxRy972"), []byte("admin"))
	// cek := helper.VerifyPassword("admin", "$2a$14$8cbHhkS0r2zwJ4br75ILpO/FqVPI.H5d1FRpwee90JhZhtVxRy972")

	// fmt.Println(err)
	// fmt.Println(cek)

	// if err := r.Run(":8080"); err != nil {
	// 	log.Fatal("Server Run Failed:", err)
	// }
}
