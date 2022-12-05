package users_controller

import (
	"encoding/json"
	"strings"

	users_service "bee-starter/src/router/users/service"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func GetUsers(ctx *beecontext.Context) {
	// Get Users
	var users = users_service.GetUsers()

	byteUsers, byteError := json.Marshal(users)
	if byteError != nil {
		panic(byteError)
	}

	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Body(byteUsers)
}

func CreateUser(ctx *beecontext.Context) {
	// Parse Body
	decoder := json.NewDecoder(ctx.Request.Body)
	var userRequest users_service.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		panic(err)
	}
	// Create New User
	id := users_service.CreateUser(userRequest)

	byteUser, byteError := json.Marshal(map[string]string{"id": id})
	if byteError != nil {
		panic(byteError)
	}

	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Body(byteUser)
}

func GetUserById(ctx *beecontext.Context) {
	var pathParams []string = strings.Split(ctx.Request.URL.Path, "/")
	var id string = pathParams[2]
	// Get User
	var user = users_service.GetUser(id)

	byteUser, byteError := json.Marshal(user)
	if byteError != nil {
		panic(byteError)
	}

	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Body(byteUser)
}

func UpdateUserById(ctx *beecontext.Context) {
	var pathParams []string = strings.Split(ctx.Request.URL.Path, "/")
	var id string = pathParams[2]
	// Parse Body
	decoder := json.NewDecoder(ctx.Request.Body)
	var userRequest users_service.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		panic(err)
	}
	// Update User
	var user = users_service.UpdateUser(id, userRequest)
	byteUser, byteError := json.Marshal(user)
	if byteError != nil {
		panic(byteError)
	}

	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Body(byteUser)
}

func PatchUserById(ctx *beecontext.Context) {
	var pathParams []string = strings.Split(ctx.Request.URL.Path, "/")
	var id string = pathParams[2]
	// Parse Body
	decoder := json.NewDecoder(ctx.Request.Body)
	var userRequest users_service.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		panic(err)
	}
	// Update User
	var user = users_service.UpdateUser(id, userRequest)
	byteUser, byteError := json.Marshal(user)
	if byteError != nil {
		panic(byteError)
	}

	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Body(byteUser)
}

func DeleteUserById(ctx *beecontext.Context) {
	var pathParams []string = strings.Split(ctx.Request.URL.Path, "/")
	var id string = pathParams[2]
	// Get User
	var user = users_service.DeleteUser(id)
	byteUser, byteError := json.Marshal(user)
	if byteError != nil {
		panic(byteError)
	}

	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Body(byteUser)
}
