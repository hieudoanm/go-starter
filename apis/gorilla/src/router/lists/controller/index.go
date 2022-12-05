package lists_controller

import (
	"encoding/json"
	"net/http"

	lists_service "gorilla-starter/src/router/lists/service"

	"github.com/gorilla/mux"
)

func GetLists(writer http.ResponseWriter, request *http.Request) {
	// Get Lists
	var lists = lists_service.GetLists()
	byteLists, err := json.Marshal(lists)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteLists)
}

func CreateList(writer http.ResponseWriter, request *http.Request) {
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var listRequest lists_service.ListRequest
	err := decoder.Decode(&listRequest)
	if err != nil {
		panic(err)
	}
	// Create New List
	id := lists_service.CreateList(listRequest)
	byteList, err := json.Marshal(map[string]string{"id": id})
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteList)
}

func GetListById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	// Get List
	var list = lists_service.GetList(id)
	byteList, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteList)
}

func UpdateListById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var listRequest lists_service.ListRequest
	err := decoder.Decode(&listRequest)
	if err != nil {
		panic(err)
	}
	// Update List
	var list = lists_service.UpdateList(id, listRequest)
	byteList, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteList)
}

func PatchListById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var listRequest lists_service.ListRequest
	err := decoder.Decode(&listRequest)
	if err != nil {
		panic(err)
	}
	// Update List
	var list = lists_service.UpdateList(id, listRequest)
	byteList, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteList)
}

func DeleteListById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	// Delete List
	var list = lists_service.DeleteList(id)
	byteList, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteList)
}
