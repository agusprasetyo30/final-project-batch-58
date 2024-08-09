package controllers

import (
	"final-project/database"
	"final-project/model"
	"final-project/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCource(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := repository.GetAllCource(database.DbConnection)

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

func GetCource(c *gin.Context) {
	var (
		result gin.H
		cource model.Cource
	)

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&cource)
	if err != nil {
		panic(err)
	}

	cource.ID = id

	data, err := repository.GetCource(database.DbConnection, cource)

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

func InsertCource(c *gin.Context) {
	cource := &model.Cource{}

	err := c.BindJSON(cource)
	if err != nil {
		panic(err)
	}

	data, err := repository.InsertCource(database.DbConnection, *cource)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, data)
}

func UpdateCource(c *gin.Context) {
	var cource model.Cource

	// Mengambil ID dari parameter
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&cource)
	if err != nil {
		panic(err)
	}

	cource.ID = id

	err = repository.UpdateCource(database.DbConnection, cource)
	if err != nil {
		panic(err)
	}

	// Digunakan untuk return data select sesuai dengan category yang dipilih
	selectCategory, err := repository.GetCource(database.DbConnection, cource)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": selectCategory,
	})
}

func DeleteCource(c *gin.Context) {
	var cource model.Cource

	// Mengambil ID dari parameter
	id, _ := strconv.Atoi(c.Param("id"))
	cource.ID = id

	err := repository.DeleteCource(database.DbConnection, cource)
	if err != nil {
		if err.Error() == "cource not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "cource not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Delete mata pelajaran success",
	})
}
