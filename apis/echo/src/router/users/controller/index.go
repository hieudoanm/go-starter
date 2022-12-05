package users_controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	users_service "echo-starter/src/router/users/service"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	// Get Users
	var users = users_service.GetUsers()
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	// Parse Body
	decoder := json.NewDecoder(c.Request().Body)
	var userRequest users_service.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		panic(err)
	}
	// Create New User
	id := users_service.CreateUser(userRequest)
	return c.JSON(http.StatusOK, map[string]string{"id": id})
}

func GetUserById(c echo.Context) error {
	var pathParams []string = strings.Split(c.Path(), "/")
	var id string = pathParams[2]
	fmt.Println(id)
	// Get User
	var user = users_service.GetUser(id)
	return c.JSON(http.StatusOK, user)
}

func UpdateUserById(c echo.Context) error {
	var pathParams []string = strings.Split(c.Path(), "")
	var id string = pathParams[2]
	// Parse Body
	decoder := json.NewDecoder(c.Request().Body)
	var userRequest users_service.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		panic(err)
	}
	// Update User
	var user = users_service.UpdateUser(id, userRequest)
	return c.JSON(http.StatusOK, user)
}

func PatchUserById(c echo.Context) error {
	var pathParams []string = strings.Split(c.Path(), "")
	var id string = pathParams[2]
	// Parse Body
	decoder := json.NewDecoder(c.Request().Body)
	var userRequest users_service.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		panic(err)
	}
	// Update User
	var user = users_service.UpdateUser(id, userRequest)
	return c.JSON(http.StatusOK, user)
}

func DeleteUserById(c echo.Context) error {
	var pathParams []string = strings.Split(c.Path(), "")
	var id string = pathParams[2]
	// Get User
	var user = users_service.DeleteUser(id)
	return c.JSON(http.StatusOK, user)
}
