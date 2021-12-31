package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

var myDB *gorm.DB

type Book struct {
	gorm.Model
	Name        string `form:"name"`
	PrintYear   int64  `form:"printYear"`
	Author      string `form:"author"`
	Description string `form:"description"`
	PageNumber  int64  `form:"pageNumber"`
}

type JsonError struct {
	Error string
}

func main() {
	connectToDB()
	serve()
}

func serve() {
	e := echo.New()
	// Routes

	e.GET("/books", AllBooks)
	e.GET("/books/:id", GetBook)

	// login
	h := &handler{}
	e.POST("/login", h.login)
	e.POST("/register", h.Register)
	var protect = middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
	})
	e.PUT("/books", CreateBook, protect)
	e.DELETE("/books/:id", DeleteBook, protect)

	// Start server
	log.Println("Server Starting on: http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}

func connectToDB() {
	dsn := "root:root123456789@tcp(0.0.0.0:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Book{}, &User{})
	if err != nil {
		panic(err)
	}
	myDB = db
}

func CreateBook(c echo.Context) error {
	var err error
	b := Book{}

	err = c.Bind(&b)
	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	if b.PageNumber == 0 {
		return c.String(http.StatusBadRequest, "page number is required")
	}

	if b.Name == "" {
		return c.String(http.StatusBadRequest, "name is required")
	}

	if b.PrintYear == 0 {
		return c.String(http.StatusBadRequest, "print year is required")
	}

	if b.Author == "" {
		return c.String(http.StatusBadRequest, "author is required")
	}

	if b.Description == "" {
		return c.String(http.StatusBadRequest, "description is required")
	}

	tx := myDB.Create(&b)
	if tx.Error != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, b)
}

func GetBook(c echo.Context) (err error) {
	idString := c.Param("id")
	var id int64
	if idString != "" {
		id, err = strconv.ParseInt(idString, 10, 32)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
	} else {
		return c.String(http.StatusBadRequest, "id is empty")
	}

	b := Book{}
	tx := myDB.Find(&b, id)
	if tx.Error != nil {
		return c.String(http.StatusBadRequest, "invalid value for page parameter")
	}

	if b.ID == 0 {
		return c.String(http.StatusNotFound, "book not found")
	}

	return c.JSON(http.StatusOK, b)
}

func DeleteBook(c echo.Context) (err error) {
	idString := c.Param("id")
	var id int64
	if idString != "" {
		id, err = strconv.ParseInt(idString, 10, 32)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
	} else {
		return c.String(http.StatusBadRequest, "id is empty")
	}

	b := Book{}
	tx := myDB.Find(&b, id)
	if tx.Error != nil {
		return c.String(http.StatusBadRequest, "invalid value for page parameter")
	}

	if b.ID == 0 {
		return c.String(http.StatusNotFound, "book not found")
	}

	tx = myDB.Delete(&b)
	if tx.Error != nil {
		return c.String(http.StatusInternalServerError, "deleted")
	}

	return c.JSON(http.StatusOK, b)
}

func AllBooks(c echo.Context) (err error) {
	pageString := c.FormValue("page")
	var page int64 = 1 // default
	if pageString != "" {
		page, err = strconv.ParseInt(pageString, 10, 32)
		if err != nil {
			return c.String(http.StatusBadRequest, "invalid value for page parameter")
		}
	}

	var books []Book

	perPage := 10
	tx := myDB.Offset(int(page-1) * perPage).Limit(10).Find(&books)
	if tx.Error != nil {
		return c.String(http.StatusInternalServerError, tx.Error.Error())
	}
	return c.JSON(http.StatusCreated, books)
}
