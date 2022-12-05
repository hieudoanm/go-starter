package users_controller

import (
	"encoding/json"
	"net/http"

	users_service "gin-starter/src/router/users/service"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	// Get Users
	var users = users_service.GetUsers()
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	// Parse Body
	decoder := json.NewDecoder(c.Request.Body)
	var userRequest users_service.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		panic(err)
	}
	// Create New User
	id := users_service.CreateUser(userRequest)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	// Get User
	var user = users_service.GetUser(id)
	c.JSON(http.StatusOK, user)
}

func UpdateUserById(c *gin.Context) {
	id := c.Param("id")
	// Parse Body
	decoder := json.NewDecoder(c.Request.Body)
	var userRequest users_service.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		panic(err)
	}
	// Update User
	var user = users_service.UpdateUser(id, userRequest)
	c.JSON(http.StatusOK, user)
}

func PatchUserById(c *gin.Context) {
	id := c.Param("id")
	// Parse Body
	decoder := json.NewDecoder(c.Request.Body)
	var userRequest users_service.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		panic(err)
	}
	// Update User
	var user = users_service.UpdateUser(id, userRequest)
	c.JSON(http.StatusOK, user)
}

func DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	// Get User
	var user = users_service.DeleteUser(id)
	c.JSON(http.StatusOK, user)
}
