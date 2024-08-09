package controllers

import (
	"final-project/database"
	"final-project/model"
	"final-project/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	var (
		result gin.H
	)

	students, err := repository.GetAllStudents(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": students,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetStudent(c *gin.Context) {
	var (
		result  gin.H
		student model.Student
	)

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&student)
	if err != nil {
		panic(err)
	}

	student.ID = id

	data, err := repository.GetStudent(database.DbConnection, student)
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

func InsertStudent(c *gin.Context) {
	student := &model.Student{}

	err := c.ShouldBindJSON(student)
	if err != nil {
		panic(err)
	}

	data, err := repository.InsertStudent(database.DbConnection, *student)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, data)
}

func UpdateStudent(c *gin.Context) {
	var student model.Student

	// Mengambil ID dari parameter
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&student)
	if err != nil {
		panic(err)
	}

	student.ID = id

	err = repository.UpdateStudent(database.DbConnection, student)
	if err != nil {
		panic(err)
	}

	// Digunakan untuk return data select sesuai dengan category yang dipilih
	selectStudent, err := repository.GetStudent(database.DbConnection, student)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": selectStudent,
	})
}

func DeleteStudent(c *gin.Context) {
	var student model.Student

	// Mengambil ID dari parameter
	id, _ := strconv.Atoi(c.Param("id"))
	student.ID = id

	err := repository.DeleteStudent(database.DbConnection, student)
	if err != nil {
		if err.Error() == "cource not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Delete student success",
	})
}
