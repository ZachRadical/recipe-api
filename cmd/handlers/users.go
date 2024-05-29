package handlers

import (
	"nomcom-api/cmd/models"
	"nomcom-api/cmd/repos"
	"net/http"
	"strconv"
	"fmt"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	newUser, err := repos.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user := models.User{}
	c.Bind(&user)
	updatedUser, err := repos.UpdateUser(user, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedUser)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user := models.User{}
	c.Bind(&user)
	desiredUser, err := repos.GetUser(user, idInt)

	if err != nil {
		fmt.Println("Error with repo method")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, desiredUser)
}

