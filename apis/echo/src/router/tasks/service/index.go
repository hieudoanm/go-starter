package tasks_service

import (
	postgres "echo-starter/src/libs/postgres"
	"echo-starter/src/utils"

	"time"

	"github.com/google/uuid"
)

var DATABASE_HOST = utils.Getenv("DATABASE_HOST", "localhost")
var DATABASE_PORT = utils.Getenv("DATABASE_PORT", "5432")
var DATABASE_USER = utils.Getenv("DATABASE_USER", "gouser")
var DATABASE_PASS = utils.Getenv("DATABASE_PASS", "gopass")
var DATABASE_NAME = utils.Getenv("DATABASE_NAME", "postgres")
var DATABASE_MODE = utils.Getenv("DATABASE_MODE", "disable")
var DATABASE_TIMEZONE = utils.Getenv("DATABASE_TIMEZONE", "Asia/Ho_Chi_Minh")

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	Createdat time.Time `json:"createdAt"`
	Updatedat time.Time `json:"updatedAt"`
}

func GetTasks() []Task {
	database := postgres.Open(postgres.DatabaseConfigs{
		Host:     DATABASE_HOST,
		Port:     DATABASE_PORT,
		User:     DATABASE_USER,
		Pass:     DATABASE_PASS,
		Name:     DATABASE_NAME,
		Mode:     DATABASE_MODE,
		TimeZone: DATABASE_TIMEZONE,
	})

	var tasks []Task
	result := database.Find(&tasks)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}

	return tasks
}

func CreateTask(title string) string {
	database := postgres.Open(postgres.DatabaseConfigs{
		Host:     DATABASE_HOST,
		Port:     DATABASE_PORT,
		User:     DATABASE_USER,
		Pass:     DATABASE_PASS,
		Name:     DATABASE_NAME,
		Mode:     DATABASE_MODE,
		TimeZone: DATABASE_TIMEZONE,
	})

	task := Task{
		ID:        uuid.New().String(),
		Title:     title,
		Completed: false,
		Createdat: time.Now(),
		Updatedat: time.Now(),
	}

	result := database.Create(&task)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}
	return task.ID
}

func GetTask(id string) Task {
	database := postgres.Open(postgres.DatabaseConfigs{
		Host:     DATABASE_HOST,
		Port:     DATABASE_PORT,
		User:     DATABASE_USER,
		Pass:     DATABASE_PASS,
		Name:     DATABASE_NAME,
		Mode:     DATABASE_MODE,
		TimeZone: DATABASE_TIMEZONE,
	})

	var task Task
	database.First(&task, "id = ?", id)

	return task
}

type UpdatedTask struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func UpdateTask(id string, updatedTask UpdatedTask) Task {
	database := postgres.Open(postgres.DatabaseConfigs{
		Host:     DATABASE_HOST,
		Port:     DATABASE_PORT,
		User:     DATABASE_USER,
		Pass:     DATABASE_PASS,
		Name:     DATABASE_NAME,
		Mode:     DATABASE_MODE,
		TimeZone: DATABASE_TIMEZONE,
	})

	var task Task = Task{ID: id}
	database.Model(&task).Updates(Task{Title: updatedTask.Title, Completed: updatedTask.Completed})

	return task
}

func DeleteTask(id string) Task {
	database := postgres.Open(postgres.DatabaseConfigs{
		Host:     DATABASE_HOST,
		Port:     DATABASE_PORT,
		User:     DATABASE_USER,
		Pass:     DATABASE_PASS,
		Name:     DATABASE_NAME,
		Mode:     DATABASE_MODE,
		TimeZone: DATABASE_TIMEZONE,
	})

	var task Task
	database.Where("id = ?", id).Delete(&task)

	return task
}
