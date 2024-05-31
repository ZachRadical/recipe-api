package handlers

import (
	"fmt"
	"net/http"
	"nomcom-api/cmd/models"
	"nomcom-api/cmd/repos"
	"nomcom-api/cmd/services"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, desiredUser)
}

func hashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashedPassword := string(hash)
	return hashedPassword, nil
}


func RegisterUser(c echo.Context) error {
	request := models.RegisterRequestBody{}
	c.Bind(&request)
	hash, err := hashPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	request.Password = hash
	newUser, err := repos.RegisterUser(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, newUser)
}

func LoginUser(c echo.Context) error {
	request := models.LoginRequestBody{}
	c.Bind(&request)

	user, err := repos.LoginUser(request)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	isValid := checkPassword(user.Password, request.Password) 
	if isValid != nil {
		fmt.Println(isValid.Error())
		return isValid
	}
	fmt.Println("ok")
	secret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := services.CreateJWT(secret, user.ID)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tokenString)

}

func checkPassword(hashedPass, plainPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plainPass))
	return err
}