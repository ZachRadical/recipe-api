package handlers

import (
	"nomcom-api/cmd/models"
	"nomcom-api/cmd/repos"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateIngredient(c echo.Context) error {
	ingredient := models.Ingredient{}
	c.Bind(&ingredient)
	newIngredient, err := repos.CreateIngredient(ingredient)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newIngredient)
}

func UpdateIngredient(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
		err.Error())
	}

	ingredient := models.Ingredient{}
	c.Bind(&ingredient)
	updatedIngredient, err := repos.UpdateIngredient(ingredient, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedIngredient)
}

func GetIngredient(c echo.Context) error {
	id := c.Param("id")
	
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
		err.Error())
	}

	ingredient := models.Ingredient{}
	c.Bind(&ingredient)
	desiredIngredient, err := repos.GetIngredient(ingredient, idInt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, desiredIngredient)
}

// func GetAllIngredientsByComponent(c echo.Context) error {
// 	id := c.Param("component_id")

// 	idInt, err := strconv.Atoi(id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError,
// 		err.Error())
// 	}

// 	ingredients := []models.Ingredient{}
// 	c.Bind(&ingredients)
// 	ingredientList, err := repos.GetAllIngredientsByComponent(idInt)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, ingredientList)
// }