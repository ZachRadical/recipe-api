package main

import (
	"github.com/labstack/echo/v4"
	"nomcom-api/cmd/handlers"
	"nomcom-api/cmd/database"
	"nomcom-api/cmd/services"

)
func main() {
	e:= echo.New()
	e.GET("/", handlers.Home)

	database.InitDB()

	e.POST("/login", handlers.LoginUser)
	e.POST("/register", handlers.RegisterUser)

	e.POST("/users", handlers.CreateUser)
	e.POST("/recipes", handlers.CreateRecipe)
	e.POST("/components", handlers.CreateComponent)
	e.POST("/ingredients", handlers.CreateIngredient)

	e.PUT("/users/:id", handlers.UpdateUser)
	e.PUT("/recipes/:id", handlers.UpdateRecipe)
	e.PUT("/components/:id", handlers.UpdateComponent)
	e.PUT("/ingredients/:id", handlers.UpdateIngredient)

	e.GET("/users/:id", handlers.GetUser)
	e.GET("/recipes/:id", handlers.GetRecipe)
	e.GET("/components/:id", handlers.GetComponent)
	e.GET("/ingredients/:id", handlers.GetIngredient)

	e.Use(services.LogRequest)
	e.Logger.Fatal(e.Start(":8080"))
}