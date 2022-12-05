package users_controller

import (
	"encoding/json"
	"net/http"

	users_service "gorilla-starter/src/router/users/service"

	"github.com/gorilla/mux"
)

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	// Get Users
	var users = users_service.GetUsers()
	byteUsers, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteUsers)
}

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var userRequest users_service.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		panic(err)
	}
	// Create New User
	id := users_service.CreateUser(userRequest)
	byteUser, err := json.Marshal(map[string]string{"id": id})
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteUser)
}

func GetUserById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	// Get User
	var user = users_service.GetUser(id)
	byteUser, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteUser)
}

func UpdateUserById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var userRequest users_service.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		panic(err)
	}
	// Update User
	var user = users_service.UpdateUser(id, userRequest)
	byteUser, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteUser)
}

func PatchUserById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var userRequest users_service.UserRequest
	err := decoder.Decode(&userRequest)
	if err != nil {
		panic(err)
	}
	// Update User
	var user = users_service.UpdateUser(id, userRequest)
	byteUser, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteUser)
}

func DeleteUserById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	// Delete User
	var user = users_service.DeleteUser(id)
	byteUser, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteUser)
}
