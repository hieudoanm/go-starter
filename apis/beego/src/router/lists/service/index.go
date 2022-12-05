package lists_service

import (
	postgres "bee-starter/src/libs/postgres"
	"bee-starter/src/utils"

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

type List struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      string    `json:"userId"`
	Createdat   time.Time `json:"createdAt"`
	Updatedat   time.Time `json:"updatedAt"`
}

type ListRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func GetLists() []List {
	database := postgres.Open(postgres.DatabaseConfigs{
		Host:     DATABASE_HOST,
		Port:     DATABASE_PORT,
		User:     DATABASE_USER,
		Pass:     DATABASE_PASS,
		Name:     DATABASE_NAME,
		Mode:     DATABASE_MODE,
		TimeZone: DATABASE_TIMEZONE,
	})

	var lists []List
	result := database.Find(&lists)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}

	return lists
}

func CreateList(listRequest ListRequest) string {
	database := postgres.Open(postgres.DatabaseConfigs{
		Host:     DATABASE_HOST,
		Port:     DATABASE_PORT,
		User:     DATABASE_USER,
		Pass:     DATABASE_PASS,
		Name:     DATABASE_NAME,
		Mode:     DATABASE_MODE,
		TimeZone: DATABASE_TIMEZONE,
	})

	list := List{
		ID:          uuid.New().String(),
		Title:       listRequest.Title,
		Description: listRequest.Description,
		Createdat:   time.Now(),
		Updatedat:   time.Now(),
	}

	result := database.Create(&list)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}
	return list.ID
}

func GetList(id string) List {
	database := postgres.Open(postgres.DatabaseConfigs{
		Host:     DATABASE_HOST,
		Port:     DATABASE_PORT,
		User:     DATABASE_USER,
		Pass:     DATABASE_PASS,
		Name:     DATABASE_NAME,
		Mode:     DATABASE_MODE,
		TimeZone: DATABASE_TIMEZONE,
	})

	var list List
	database.First(&list, "id = ?", id)

	return list
}

type UpdatedList struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func UpdateList(id string, listRequest ListRequest) List {
	database := postgres.Open(postgres.DatabaseConfigs{
		Host:     DATABASE_HOST,
		Port:     DATABASE_PORT,
		User:     DATABASE_USER,
		Pass:     DATABASE_PASS,
		Name:     DATABASE_NAME,
		Mode:     DATABASE_MODE,
		TimeZone: DATABASE_TIMEZONE,
	})

	var list List = List{ID: id}
	database.Model(&list).Updates(List{
		Title:       listRequest.Title,
		Description: listRequest.Description,
	})

	return list
}

func DeleteList(id string) List {
	database := postgres.Open(postgres.DatabaseConfigs{
		Host:     DATABASE_HOST,
		Port:     DATABASE_PORT,
		User:     DATABASE_USER,
		Pass:     DATABASE_PASS,
		Name:     DATABASE_NAME,
		Mode:     DATABASE_MODE,
		TimeZone: DATABASE_TIMEZONE,
	})

	var list List
	database.Where("id = ?", id).Delete(&list)

	return list
}
