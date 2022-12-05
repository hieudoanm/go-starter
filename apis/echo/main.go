// Echo Starter:
//
//	version: 0.0.1
//	description: Echo Starter
//
// Schemes: http, https
// BasePath: /
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//
// swagger:meta
package main

import (
	health_controller "echo-starter/src/router/health/controller"
	lists_controller "echo-starter/src/router/lists/controller"
	tasks_controller "echo-starter/src/router/tasks/controller"
	users_controller "echo-starter/src/router/users/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/health", health_controller.GetHealth)
	// Tasks
	e.GET("/tasks", tasks_controller.GetTasks)
	e.POST("/tasks", tasks_controller.CreateTask)
	e.GET("/tasks/:id", tasks_controller.GetTaskById)
	e.PUT("/tasks/:id", tasks_controller.UpdateTaskById)
	e.PATCH("/tasks/:id", tasks_controller.UpdateTaskById)
	e.DELETE("/tasks/:id", tasks_controller.DeleteTaskById)
	// Lists
	e.GET("/lists", lists_controller.GetLists)
	e.POST("/lists", lists_controller.CreateList)
	e.GET("/lists/:id", lists_controller.GetListById)
	e.PUT("/lists/:id", lists_controller.UpdateListById)
	e.PATCH("/lists/:id", lists_controller.UpdateListById)
	e.DELETE("/lists/:id", lists_controller.DeleteListById)
	// Users
	e.GET("/users", users_controller.GetUsers)
	e.POST("/users", users_controller.CreateUser)
	e.GET("/users/:id", users_controller.GetUserById)
	e.PUT("/users/:id", users_controller.UpdateUserById)
	e.PATCH("/users/:id", users_controller.UpdateUserById)
	e.DELETE("/users/:id", users_controller.DeleteUserById)

	e.Logger.Fatal(e.Start(":8080"))
}
