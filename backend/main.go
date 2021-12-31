package main

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"

	"gorm.io/driver/mysql"
)

var myDB *gorm.DB

type Book struct {
	gorm.Model
	Name        string
	PublishYear int64
	Author      string
	Description string
	Pages       int64
}

type User struct {
	gorm.Model
	UserName string
	Email    string
	Password string
}

func main() {
	connectToDB()
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/create", CreateBook)
	http.HandleFunc("/get", GetBook)
	panic(http.ListenAndServe(":8080", nil))
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

func CreateUser() {

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == "PUT" {
		b := Book{
			Name:        r.FormValue("name"),
			Author:      r.FormValue("author"),
			Description: r.FormValue("description"),
		}

		if r.FormValue("PublishYear") != "" {
			b.PublishYear, err = strconv.ParseInt(r.FormValue("PublishYear"), 10, 32)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
		}

		if r.FormValue("pages") != "" {
			b.Pages, err = strconv.ParseInt(r.FormValue("pages"), 10, 32)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
		}

		tx := myDB.Create(&b)
		if tx.Error != nil {
			http.Error(w, tx.Error.Error(), 400)
			return
		}

		jData, err := json.Marshal(b)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jData)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	} else {
		http.Redirect(w, r, "/", 301)
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != "GET" {
		http.Redirect(w, r, "/", 301)
		return
	}

	idString := r.FormValue("id")
	var id int64
	if idString != "" {
		id, err = strconv.ParseInt(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		http.Error(w, "id is empty", 400)
		return
	}

	b := Book{}
	tx := myDB.Find(&b, id)
	if tx.Error != nil {
		http.Error(w, tx.Error.Error(), 400)
		return
	}

	if b.ID == 0 {
		http.Error(w, "book not found", 404)
		return
	}

	jData, err := json.Marshal(b)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jData)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
