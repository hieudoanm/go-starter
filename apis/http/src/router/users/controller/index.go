package users_controller

import (
	"encoding/json"
	"net/http"

	users_service "http-starter/src/router/users/service"

	"github.com/julienschmidt/httprouter"
)

func GetUsers(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	// Get Users
	var users = users_service.GetUsers()
	json.NewEncoder(writer).Encode(users)
}

func CreateUser(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var userRequest users_service.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		panic(err)
	}
	// Create New User
	id := users_service.CreateUser(userRequest)
	json.NewEncoder(writer).Encode(map[string]string{"id": id})
}

func GetUserById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var id string = params.ByName("id")
	// Get User
	var user = users_service.GetUser(id)
	json.NewEncoder(writer).Encode(user)
}

func UpdateUserById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var id string = params.ByName("id")
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var userRequest users_service.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		panic(err)
	}
	// Update User
	var user = users_service.UpdateUser(id, userRequest)
	json.NewEncoder(writer).Encode(user)
}

func PatchUserById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var id string = params.ByName("id")
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var userRequest users_service.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		panic(err)
	}
	// Update User
	var user = users_service.UpdateUser(id, userRequest)
	json.NewEncoder(writer).Encode(user)
}

func DeleteUserById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var id string = params.ByName("id")
	// Get User
	var user = users_service.DeleteUser(id)
	json.NewEncoder(writer).Encode(user)
}
