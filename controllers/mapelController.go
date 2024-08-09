package controllers

import (
	"final-project/database"
	"final-project/model"
	"final-project/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMapel(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := repository.GetAllMapel(database.DbConnection)

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

func GetMapel(c *gin.Context) {
	var (
		result gin.H
		mapel  model.Mapel
	)

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&mapel)
	if err != nil {
		panic(err)
	}

	mapel.ID = id

	data, err := repository.GetMapel(database.DbConnection, mapel)

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

func InsertMapel(c *gin.Context) {
	mapel := &model.Mapel{}

	err := c.BindJSON(mapel)
	if err != nil {
		panic(err)
	}

	data, err := repository.InsertMapel(database.DbConnection, *mapel)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, data)
}

func UpdateMapel(c *gin.Context) {
	var mapel model.Mapel

	// Mengambil ID dari parameter
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&mapel)
	if err != nil {
		panic(err)
	}

	mapel.ID = id

	err = repository.UpdateMapel(database.DbConnection, mapel)
	if err != nil {
		panic(err)
	}

	// Digunakan untuk return data select sesuai dengan category yang dipilih
	selectCategory, err := repository.GetMapel(database.DbConnection, mapel)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": selectCategory,
	})
}

func DeleteMapel(c *gin.Context) {
	var mapel model.Mapel

	// Mengambil ID dari parameter
	id, _ := strconv.Atoi(c.Param("id"))
	mapel.ID = id

	err := repository.DeleteMapel(database.DbConnection, mapel)
	if err != nil {
		if err.Error() == "mapel not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "mapel not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Delete mata pelajaran success",
	})
}
