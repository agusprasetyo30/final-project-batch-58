package controllers

import (
	"final-project/database"
	"final-project/model"
	"final-project/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllAssessment(c *gin.Context) {
	var (
		result gin.H
	)

	assessments, err := repository.GetAllAssessment(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": assessments,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetAssessment(c *gin.Context) {
	var (
		result     gin.H
		assessment model.Assessment
	)

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&assessment)
	if err != nil {
		panic(err)
	}

	assessment.ID = id

	data, err := repository.GetAssessment(database.DbConnection, assessment)
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

func InsertAssessment(c *gin.Context) {
	assessment := &model.Assessment{}

	err := c.BindJSON(assessment)
	if err != nil {
		panic(err)
	}

	// konversi
	assessment.Grade = gradeConversion(assessment.Score)

	data, err := repository.InsertAssessment(database.DbConnection, *assessment)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, data)
}

func UpdateAssessment(c *gin.Context) {
	var assessment model.Assessment

	// Mengambil ID dari parameter
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&assessment)
	if err != nil {
		panic(err)
	}

	assessment.ID = id

	// konversi
	assessment.Grade = gradeConversion(assessment.Score)

	err = repository.UpdateAssessment(database.DbConnection, assessment)
	if err != nil {
		panic(err)
	}

	// Digunakan untuk return data select sesuai dengan category yang dipilih
	selectAssessment, err := repository.GetAssessment(database.DbConnection, assessment)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": selectAssessment,
	})
}

func DeleteAssessment(c *gin.Context) {
	var assessment model.Assessment

	// Mengambil ID dari parameter
	id, _ := strconv.Atoi(c.Param("id"))
	assessment.ID = id

	err := repository.DeleteAssessment(database.DbConnection, assessment)
	if err != nil {
		if err.Error() == "cource not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "assessment not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Delete assessment success",
	})
}

func gradeConversion(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "E"
	}
}
