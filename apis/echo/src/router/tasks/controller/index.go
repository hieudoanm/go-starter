package tasks_controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	tasks_service "echo-starter/src/router/tasks/service"

	"github.com/labstack/echo/v4"
)

func GetTasks(c echo.Context) error {
	// Get Tasks
	var tasks = tasks_service.GetTasks()
	return c.JSON(http.StatusOK, tasks)
}

type NewTask struct {
	Title string `json:"title"`
}

func CreateTask(c echo.Context) error {
	// Parse Body
	decoder := json.NewDecoder(c.Request().Body)
	var newTask NewTask
	err := decoder.Decode(&newTask)
	if err != nil {
		panic(err)
	}
	// Create New Task
	id := tasks_service.CreateTask(newTask.Title)
	return c.JSON(http.StatusOK, map[string]string{"id": id})
}

func GetTaskById(c echo.Context) error {
	var pathParams []string = strings.Split(c.Path(), "/")
	var id string = pathParams[2]
	fmt.Println(id)
	// Get Task
	var task = tasks_service.GetTask(id)
	return c.JSON(http.StatusOK, task)
}

func UpdateTaskById(c echo.Context) error {
	var pathParams []string = strings.Split(c.Path(), "")
	var id string = pathParams[2]
	// Parse Body
	decoder := json.NewDecoder(c.Request().Body)
	var updatedTask tasks_service.UpdatedTask
	err := decoder.Decode(&updatedTask)
	if err != nil {
		panic(err)
	}
	// Update Task
	var task = tasks_service.UpdateTask(id, updatedTask)
	return c.JSON(http.StatusOK, task)
}

func PatchTaskById(c echo.Context) error {
	var pathParams []string = strings.Split(c.Path(), "")
	var id string = pathParams[2]
	// Parse Body
	decoder := json.NewDecoder(c.Request().Body)
	var updatedTask tasks_service.UpdatedTask
	err := decoder.Decode(&updatedTask)
	if err != nil {
		panic(err)
	}
	// Update Task
	var task = tasks_service.UpdateTask(id, updatedTask)
	return c.JSON(http.StatusOK, task)
}

func DeleteTaskById(c echo.Context) error {
	var pathParams []string = strings.Split(c.Path(), "")
	var id string = pathParams[2]
	// Get Task
	var task = tasks_service.DeleteTask(id)
	return c.JSON(http.StatusOK, task)
}
