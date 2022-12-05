package lists_service

import (
	postgres "gorilla-starter/src/libs/postgres"

	"time"

	"github.com/google/uuid"
)

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
	UserId      string `json:"userId"`
}

func GetLists() []List {
	database := postgres.OpenDatabase()

	var lists []List
	result := database.Find(&lists)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}

	return lists
}

func CreateList(listRequest ListRequest) string {
	database := postgres.OpenDatabase()

	list := List{
		ID:          uuid.New().String(),
		Title:       listRequest.Title,
		Description: listRequest.Description,
		UserId:      listRequest.UserId,
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
	database := postgres.OpenDatabase()

	var list List
	database.First(&list, "id = ?", id)

	return list
}

func UpdateList(id string, listRequest ListRequest) List {
	database := postgres.OpenDatabase()

	var list List = List{ID: id}
	database.Model(&list).Updates(List{
		Title:       listRequest.Title,
		Description: listRequest.Description,
		UserId:      listRequest.UserId,
	})

	return list
}

func DeleteList(id string) List {
	database := postgres.OpenDatabase()

	var list List
	database.Where("id = ?", id).Delete(&list)

	return list
}
