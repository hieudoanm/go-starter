package lists_controller

import (
	"encoding/json"
	"net/http"

	lists_service "http-starter/src/router/lists/service"

	"github.com/julienschmidt/httprouter"
)

func GetLists(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	// Get Lists
	var lists = lists_service.GetLists()
	json.NewEncoder(writer).Encode(lists)
}

func CreateList(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var listRequest lists_service.ListRequest
	err := decoder.Decode(&listRequest)
	if err != nil {
		panic(err)
	}
	// Create New List
	id := lists_service.CreateList(listRequest)
	json.NewEncoder(writer).Encode(map[string]string{"id": id})
}

func GetListById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var id string = params.ByName("id")
	// Get List
	var list = lists_service.GetList(id)
	json.NewEncoder(writer).Encode(list)
}

func UpdateListById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var id string = params.ByName("id")
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var listRequest lists_service.ListRequest
	err := decoder.Decode(&listRequest)
	if err != nil {
		panic(err)
	}
	// Update List
	var list = lists_service.UpdateList(id, listRequest)
	json.NewEncoder(writer).Encode(list)
}

func PatchListById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var id string = params.ByName("id")
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var listRequest lists_service.ListRequest
	err := decoder.Decode(&listRequest)
	if err != nil {
		panic(err)
	}
	// Update List
	var list = lists_service.UpdateList(id, listRequest)
	json.NewEncoder(writer).Encode(list)
}

func DeleteListById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var id string = params.ByName("id")
	// Get List
	var list = lists_service.DeleteList(id)
	json.NewEncoder(writer).Encode(list)
}
