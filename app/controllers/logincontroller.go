package controllers

import (
	"net/http"
	"time"

	"github.com/alghiffaryfa19/echo-rest/app/models"
	"github.com/alghiffaryfa19/echo-rest/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CheckLogin(c echo.Context) error {
	email := c.FormValue(("email"))
	password := c.FormValue(("password"))

	res, err := models.CheckLogin(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages" : err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages" : err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})

}

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")
	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

