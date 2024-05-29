package handlers

import (
	"nomcom-api/cmd/models"
	"nomcom-api/cmd/repos"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


func CreateProduct(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)
	newProduct, err := repos.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newProduct)
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
		err.Error())
	}

	product := models.Product{}
	c.Bind(&product)
	updatedProduct, err := repos.UpdateProduct(product, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedProduct)
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
		err.Error())
	}

	product := models.Product{}
	c.Bind(&product)
	desiredProduct, err := repos.GetProduct(product, idInt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, desiredProduct)
}

// func GetAllProductsByUser(c echo.Context) error {
// 	id := c.Param("user_id")

// 	idInt, err := strconv.Atoi(id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError,
// 		err.Error())
// 	}

// 	products := []models.Product{}
// 	c.Bind(&products)
// 	productList, err := repos.GetAllProductsByUser(idInt)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, productList)
// }