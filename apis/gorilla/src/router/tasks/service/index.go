package tasks_service

import (
	postgres "gorilla-starter/src/libs/postgres"

	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	ListId      string    `json:"listId"`
	Createdat   time.Time `json:"createdAt"`
	Updatedat   time.Time `json:"updatedAt"`
}

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	ListId      string `json:"listId"`
}

func GetTasks() []Task {
	database := postgres.OpenDatabase()

	var tasks []Task
	result := database.Find(&tasks)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}

	return tasks
}

func CreateTask(taskRequest TaskRequest) string {
	database := postgres.OpenDatabase()

	task := Task{
		ID:          uuid.New().String(),
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
		Completed:   taskRequest.Completed,
		ListId:      taskRequest.ListId,
		Createdat:   time.Now(),
		Updatedat:   time.Now(),
	}

	result := database.Create(&task)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}
	return task.ID
}

func GetTask(id string) Task {
	database := postgres.OpenDatabase()

	var task Task
	database.First(&task, "id = ?", id)

	return task
}

func UpdateTask(id string, taskRequest TaskRequest) Task {
	database := postgres.OpenDatabase()

	var task Task = Task{ID: id}
	database.Model(&task).Updates(Task{
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
		Completed:   taskRequest.Completed,
		ListId:      taskRequest.ListId,
	})

	return task
}

func DeleteTask(id string) Task {
	database := postgres.OpenDatabase()

	var task Task
	database.Where("id = ?", id).Delete(&task)

	return task
}
