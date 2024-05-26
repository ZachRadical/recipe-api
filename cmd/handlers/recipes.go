package handlers

import (
	"nomcom-api/cmd/models"
	"nomcom-api/cmd/repos"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateRecipe(c echo.Context) error {
	recipe := models.Recipe{}
	c.Bind(&recipe)
	newRecipe, err := repos.CreateRecipe(recipe)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newRecipe)
}