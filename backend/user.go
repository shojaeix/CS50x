package main

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	UserName string `json:"user_name" gorm:"unique"`
	Email    string
	Password string
}

type handler struct{}

type LoginParameters struct {
	Username string `json:"user_name"`
	Password string
}

func (h *handler) login(c echo.Context) (err error) {
	params := &LoginParameters{}
	err = c.Bind(params)
	if err != nil {
		return err
	}
	var user User
	if params.Username == "" || params.Password == "" {
		return c.JSON(http.StatusBadRequest, "username and password required")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	tx := myDB.Where("user_name", params.Username).Find(&user)
	if tx.Error != nil {
		return tx.Error
	}

	if user.ID == 0 {

		return echo.ErrUnauthorized

	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(params.Password))
	if err != nil {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID          works too)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})

}

func (h *handler) Register(c echo.Context) error {
	var err error
	u := User{}

	err = c.Bind(&u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if u.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "password is required")

	}

	if u.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name is required")
	}

	if u.UserName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "username is required")
	}

	if u.Email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "email is required")
	}
	count := int64(0)
	tx := myDB.Model(&User{}).Where("user_name", u.UserName).Count(&count)
	if tx.Error != nil {
		return tx.Error
	}
	if count > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "username already taken")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	tx = myDB.Create(&u)
	if tx.Error != nil {
		return echo.ErrInternalServerError
	}

	u.Password = ""
	return c.JSON(http.StatusOK, u)
}
