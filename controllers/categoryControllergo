package controllers

import (
	"net/http"
	"quiz-3/database"
	"quiz-3/model"
	"quiz-3/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := repository.GetAllCategory(database.DbConnection)

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

func GetCategory(c *gin.Context) {
	var (
		result   gin.H
		category model.Category
	)

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.ID = id

	category1, err := repository.GetCategory(database.DbConnection, category)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": category1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	category := &model.Category{}

	err := c.BindJSON(category)
	if err != nil {
		panic(err)
	}

	// Mengambil data user setelah login
	user, ok := c.Get("user")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// digunakan untuk mengambil username dan ditambahkan ke created_by
	category.CreatedBy = user.(*model.User).Username

	err = repository.InsertCategory(database.DbConnection, *category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, category)
}

func UpdateCategory(c *gin.Context) {
	var category model.Category

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&category)
	if err != nil {
		panic(err)
	}

	// Mengambil data user setelah login
	user, ok := c.Get("user")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	category.ID = id
	// digunakan untuk mengambil username dan ditambahkan ke created_by
	category.ModifiedBy = &user.(*model.User).Username

	err = repository.UpdateCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	// Digunakan untuk return data select sesuai dengan category yang dipilih
	selectCategory, err := repository.GetCategory(database.DbConnection, category)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, selectCategory)
}

func DeleteCategory(c *gin.Context) {
	var category model.Category

	id, _ := strconv.Atoi(c.Param("id"))

	category.ID = id

	err := repository.DeleteCategory(database.DbConnection, category)
	if err != nil {
		if err.Error() == "category not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Delete category success",
	})
}

func GeBookByCategory(c *gin.Context) {
	var result gin.H

	id, _ := strconv.Atoi(c.Param("id"))

	book1, err := repository.GetBookByCategory(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": book1,
		}
	}

	c.JSON(http.StatusOK, result)
}
