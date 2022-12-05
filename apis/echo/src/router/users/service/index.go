package users_service

import (
	postgres "echo-starter/src/libs/postgres"

	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Createdat time.Time `json:"createdAt"`
	Updatedat time.Time `json:"updatedAt"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUsers() []User {
	database := postgres.OpenDatabase()

	var users []User
	result := database.Find(&users)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}

	return users
}

func CreateUser(userRequest UserRequest) string {
	database := postgres.OpenDatabase()

	user := User{
		ID:        uuid.New().String(),
		Username:  userRequest.Username,
		Password:  userRequest.Password,
		Createdat: time.Now(),
		Updatedat: time.Now(),
	}

	result := database.Create(&user)
	resultError := result.Error
	if resultError != nil {
		panic(resultError)
	}
	return user.ID
}

func GetUser(id string) User {
	database := postgres.OpenDatabase()

	var user User
	database.First(&user, "id = ?", id)

	return user
}

func UpdateUser(id string, userRequest UserRequest) User {
	database := postgres.OpenDatabase()

	var user User = User{ID: id}
	database.Model(&user).Updates(User{
		Username: userRequest.Username,
		Password: userRequest.Password,
	})

	return user
}

func DeleteUser(id string) User {
	database := postgres.OpenDatabase()

	var user User
	database.Where("id = ?", id).Delete(&user)

	return user
}
