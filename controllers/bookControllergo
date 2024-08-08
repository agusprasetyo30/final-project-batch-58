package controllers

import (
	"net/http"
	"quiz-3/database"
	"quiz-3/model"
	"quiz-3/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBook(c *gin.Context) {
	var (
		result gin.H
	)

	books, err := repository.GetAllBook(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	book := &model.Book{}

	err := c.BindJSON(book)
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
	book.CreatedBy = user.(*model.User).Username
	book.Category = model.Category{}

	// Validation
	if book.ReleaseYear <= 1980 || book.ReleaseYear >= 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid year input"})
		return
	}

	// Konversi
	if book.TotalPage >= 100 {
		book.Thickness = "Tebal"
	} else {
		book.Thickness = "Tipis"
	}

	err = repository.InsertBook(database.DbConnection, *book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, book)
}

func GetBook(c *gin.Context) {
	var (
		result gin.H
		book   model.Book
	)

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&book)
	if err != nil {
		panic(err)
	}

	book.ID = id

	book1, err := repository.GetBook(database.DbConnection, book)

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

func UpdateBook(c *gin.Context) {
	var book model.Book

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&book)
	if err != nil {
		panic(err)
	}

	// Mengambil data user setelah login
	user, ok := c.Get("user")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	book.ID = id
	// digunakan untuk mengambil username dan ditambahkan ke created_by
	book.ModifiedBy = &user.(*model.User).Username

	// Validation
	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid year input"})
		return
	}

	// Konversi
	if book.TotalPage >= 100 {
		book.Thickness = "Tebal"
	} else {
		book.Thickness = "Tipis"
	}

	err = repository.UpdateBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	// Digunakan untuk return data select sesuai dengan category yang dipilih
	selectBook, err := repository.GetBook(database.DbConnection, book)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, selectBook)
}

func DeleteBook(c *gin.Context) {
	var book model.Book

	id, _ := strconv.Atoi(c.Param("id"))

	book.ID = id

	err := repository.DeleteBook(database.DbConnection, book)
	if err != nil {
		if err.Error() == "Book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Delete Book success",
	})
}
