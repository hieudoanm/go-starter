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

func CreateTask(c *gin.Context) {
	// Parse Body
	decoder := json.NewDecoder(c.Request.Body)
	var taskRequest tasks_service.TaskRequest
	err := decoder.Decode(&taskRequest)
	if err != nil {
		panic(err)
	}
	// Create New Task
	id := tasks_service.CreateTask(taskRequest)
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
	var taskRequest tasks_service.TaskRequest
	err := decoder.Decode(&taskRequest)
	if err != nil {
		panic(err)
	}
	// Update Task
	var task = tasks_service.UpdateTask(id, taskRequest)
	c.JSON(http.StatusOK, task)
}

func PatchTaskById(c *gin.Context) {
	id := c.Param("id")
	// Parse Body
	decoder := json.NewDecoder(c.Request.Body)
	var taskRequest tasks_service.TaskRequest
	err := decoder.Decode(&taskRequest)
	if err != nil {
		panic(err)
	}
	// Update Task
	var task = tasks_service.UpdateTask(id, taskRequest)
	c.JSON(http.StatusOK, task)
}

func DeleteTaskById(c *gin.Context) {
	id := c.Param("id")
	// Get Task
	var task = tasks_service.DeleteTask(id)
	c.JSON(http.StatusOK, task)
}
