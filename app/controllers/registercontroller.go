package controllers

import (
	"net/http"
	"strconv"

	"github.com/alghiffaryfa19/echo-rest/app/models"
	"github.com/alghiffaryfa19/echo-rest/helpers"
	"github.com/labstack/echo"
)


func Register(c echo.Context) error{
	name  := c.FormValue("name")
	email  := c.FormValue("email")
	password  := c.FormValue("password")
	role  := c.FormValue("role")

	hash, err := helpers.HashPassword(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	role_id, err := strconv.Atoi(role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.Register(name, email, hash, role_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}