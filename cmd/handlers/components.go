package handlers

import (
	
	"net/http"
	"nomcom-api/cmd/models"
	"nomcom-api/cmd/repos"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateComponent(c echo.Context) error {
	component := models.Component{}
	c.Bind(&component)
	newComponent, err := repos.CreateComponent(component)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newComponent)
}

func UpdateComponent(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
		err.Error())
	}

	component := models.Component{}
	c.Bind(&component)
	updatedComponent, err := repos.UpdateComponent(component, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedComponent)
}

func GetComponent(c echo.Context) error {
	id := c.Param("id")
	
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
		err.Error())
	}

	component := models.Component{}
	c.Bind(&component)
	desiredComponent, err := repos.GetComponent(component, idInt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, desiredComponent)
}

// func GetAllComponentsByRecipe(c echo.Context) error {
// 	id := c.Param("recipe_id")

// 	idInt, err := strconv.Atoi(id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError,
// 		err.Error())
// 	}

// 	components := []models.Component{}
// 	c.Bind(&components)
// 	componentList, err := repos.GetAllComponentsByRecipe(idInt)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, componentList)
// }