package tasks_controller

import (
	"encoding/json"
	"net/http"

	tasks_service "gorilla-starter/src/router/tasks/service"

	"github.com/gorilla/mux"
)

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	// Get Tasks
	var tasks = tasks_service.GetTasks()
	byteTasks, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteTasks)
}

type NewTask struct {
	Title string `json:"title"`
}

func CreateTask(writer http.ResponseWriter, request *http.Request) {
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var newTask NewTask
	err := decoder.Decode(&newTask)
	if err != nil {
		panic(err)
	}
	// Create New Task
	id := tasks_service.CreateTask(newTask.Title)
	byteTask, err := json.Marshal(map[string]string{"id": id})
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteTask)
}

func GetTaskById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	// Get Task
	var task = tasks_service.GetTask(id)
	byteTask, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteTask)
}

func UpdateTaskById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var updatedTask tasks_service.UpdatedTask
	err := decoder.Decode(&updatedTask)
	if err != nil {
		panic(err)
	}
	// Update Task
	var task = tasks_service.UpdateTask(id, updatedTask)
	byteTask, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteTask)
}

func PatchTaskById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	// Parse Body
	decoder := json.NewDecoder(request.Body)
	var updatedTask tasks_service.UpdatedTask
	err := decoder.Decode(&updatedTask)
	if err != nil {
		panic(err)
	}
	// Update Task
	var task = tasks_service.UpdateTask(id, updatedTask)
	byteTask, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteTask)
}

func DeleteTaskById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	// Delete Task
	var task = tasks_service.DeleteTask(id)
	byteTask, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteTask)
}
