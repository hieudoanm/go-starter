package tasks_controller

import (
	"encoding/json"
	"net/http"

	tasks_service "http-starter/src/router/tasks/service"

	"github.com/julienschmidt/httprouter"
)

func GetTasks(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	// Get Tasks
	var tasks = tasks_service.GetTasks()
	json.NewEncoder(writer).Encode(tasks)
}

type NewTask struct {
	Title string `json:"title"`
}

func CreateTask(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var newTask NewTask
	err := decoder.Decode(&newTask)
	if err != nil {
		panic(err)
	}
	// Create New Task
	id := tasks_service.CreateTask(newTask.Title)
	json.NewEncoder(writer).Encode(map[string]string{"id": id})
}

func GetTaskById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var id string = params.ByName("id")
	// Get Task
	var task = tasks_service.GetTask(id)
	json.NewEncoder(writer).Encode(task)
}

func UpdateTaskById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var id string = params.ByName("id")
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var updatedTask tasks_service.UpdatedTask
	err := decoder.Decode(&updatedTask)
	if err != nil {
		panic(err)
	}
	// Update Task
	var task = tasks_service.UpdateTask(id, updatedTask)
	json.NewEncoder(writer).Encode(task)
}

func PatchTaskById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var id string = params.ByName("id")
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var updatedTask tasks_service.UpdatedTask
	err := decoder.Decode(&updatedTask)
	if err != nil {
		panic(err)
	}
	// Update Task
	var task = tasks_service.UpdateTask(id, updatedTask)
	json.NewEncoder(writer).Encode(task)
}

func DeleteTaskById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var id string = params.ByName("id")
	// Get Task
	var task = tasks_service.DeleteTask(id)
	json.NewEncoder(writer).Encode(task)
}
