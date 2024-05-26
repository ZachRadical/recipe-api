package main

import (
	"github.com/labstack/echo/v4"
	"nomcom-api/cmd/handlers"
	"nomcom-api/cmd/database"
)
func main() {
	e:= echo.New()
	e.GET("/", handlers.Home)

	database.InitDB()

	e.POST("/users", handlers.CreateUser)
	e.POST("/recipes", handlers.CreateRecipe)
	e.Logger.Fatal(e.Start(":8080"))
}