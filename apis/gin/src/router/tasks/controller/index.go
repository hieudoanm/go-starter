package tasks_controller

import (
	"encoding/json"
	"net/http"

	tasks_service "gin-starter/src/router/tasks/service"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	// Get Tasks
	var tasks = tasks_service.GetTasks()
	c.JSON(http.StatusOK, tasks)
}

type NewTask struct {
	Title string `json:"title"`
}

func CreateTask(c *gin.Context) {
	// Parse Body
	decoder := json.NewDecoder(c.Request.Body)
	var newTask NewTask
	err := decoder.Decode(&newTask)
	if err != nil {
		panic(err)
	}
	// Create New Task
	id := tasks_service.CreateTask(newTask.Title)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	// Get Task
	var task = tasks_service.GetTask(id)
	c.JSON(http.StatusOK, task)
}

func UpdateTaskById(c *gin.Context) {
	id := c.Param("id")
	// Parse Body
	decoder := json.NewDecoder(c.Request.Body)
	var updatedTask tasks_service.UpdatedTask
	err := decoder.Decode(&updatedTask)
	if err != nil {
		panic(err)
	}
	// Update Task
	var task = tasks_service.UpdateTask(id, updatedTask)
	c.JSON(http.StatusOK, task)
}

func PatchTaskById(c *gin.Context) {
	id := c.Param("id")
	// Parse Body
	decoder := json.NewDecoder(c.Request.Body)
	var updatedTask tasks_service.UpdatedTask
	err := decoder.Decode(&updatedTask)
	if err != nil {
		panic(err)
	}
	// Update Task
	var task = tasks_service.UpdateTask(id, updatedTask)
	c.JSON(http.StatusOK, task)
}

func DeleteTaskById(c *gin.Context) {
	id := c.Param("id")
	// Get Task
	var task = tasks_service.DeleteTask(id)
	c.JSON(http.StatusOK, task)
}
