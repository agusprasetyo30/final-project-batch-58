package controllers

import (
	"final-project/database"
	"final-project/model"
	"final-project/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllClasses(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := repository.GetAllClasses(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetClass(c *gin.Context) {
	var (
		result gin.H
		class  model.Class
	)

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&class)
	if err != nil {
		panic(err)
	}

	class.ID = id

	data, err := repository.GetClass(database.DbConnection, class)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": data,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertClass(c *gin.Context) {
	class := &model.Class{}

	err := c.BindJSON(class)
	if err != nil {
		panic(err)
	}

	// Mengambil data user setelah login
	// user, ok := c.Get("user")

	// if !ok {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	// 	return
	// }

	// digunakan untuk mengambil username dan ditambahkan ke created_by
	// category.CreatedBy = user.(*model.User).Username

	data, err := repository.InsertClass(database.DbConnection, *class)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, data)
}

func UpdateClass(c *gin.Context) {
	var class model.Class

	// Mengambil ID dari parameter
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&class)
	if err != nil {
		panic(err)
	}

	// Mengambil data user setelah login
	// user, ok := c.Get("user")
	// if !ok {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	// 	return
	// }

	class.ID = id

	err = repository.UpdateClass(database.DbConnection, class)
	if err != nil {
		panic(err)
	}

	// Digunakan untuk return data select sesuai dengan category yang dipilih
	selectCategory, err := repository.GetClass(database.DbConnection, class)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": selectCategory,
	})
}

func DeleteClass(c *gin.Context) {
	var class model.Class

	// Mengambil ID dari parameter
	id, _ := strconv.Atoi(c.Param("id"))
	class.ID = id

	err := repository.DeleteClass(database.DbConnection, class)
	if err != nil {
		if err.Error() == "class not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "class not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Delete class success",
	})
}
