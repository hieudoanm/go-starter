// Gin Starter:
//
//	version: 0.0.1
//	title: Gin Starter
//
// Schemes: http, https
// Host: localhost:8080
// BasePath: /
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	swagger:meta
package main

import (
	health_controller "gin-starter/src/router/health/controller"
	lists_controller "gin-starter/src/router/lists/controller"
	tasks_controller "gin-starter/src/router/tasks/controller"
	users_controller "gin-starter/src/router/users/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", health_controller.GetHealth)
	// Tasks
	r.GET("/tasks", tasks_controller.GetTasks)
	r.POST("/tasks", tasks_controller.CreateTask)
	r.GET("/tasks/:id", tasks_controller.GetTaskById)
	r.PUT("/tasks/:id", tasks_controller.UpdateTaskById)
	r.PATCH("/tasks/:id", tasks_controller.PatchTaskById)
	r.DELETE("/tasks/:id", tasks_controller.DeleteTaskById)
	// Lists
	r.GET("/lists", lists_controller.GetLists)
	r.POST("/lists", lists_controller.CreateList)
	r.GET("/lists/:id", lists_controller.GetListById)
	r.PUT("/lists/:id", lists_controller.UpdateListById)
	r.PATCH("/lists/:id", lists_controller.PatchListById)
	r.DELETE("/lists/:id", lists_controller.DeleteListById)
	// Users
	r.GET("/users", users_controller.GetUsers)
	r.POST("/users", users_controller.CreateUser)
	r.GET("/users/:id", users_controller.GetUserById)
	r.PUT("/users/:id", users_controller.UpdateUserById)
	r.PATCH("/users/:id", users_controller.PatchUserById)
	r.DELETE("/users/:id", users_controller.DeleteUserById)

	r.Run()
}
