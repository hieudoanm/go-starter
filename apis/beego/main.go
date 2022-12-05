// Beego Starter:
//
//	version: 0.0.1
//	description: Beego Starter
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
	health_controller "bee-starter/src/router/health/controller"
	lists_controller "bee-starter/src/router/lists/controller"
	tasks_controller "bee-starter/src/router/tasks/controller"
	users_controller "bee-starter/src/router/users/controller"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Get("/health", health_controller.GetHealth)
	// List
	web.Get("/lists", lists_controller.GetLists)
	web.Post("/lists", lists_controller.CreateList)
	web.Get("/lists/:id", lists_controller.GetListById)
	web.Put("/lists/:id", lists_controller.UpdateListById)
	web.Patch("/lists/:id", lists_controller.UpdateListById)
	web.Delete("/lists/:id", lists_controller.DeleteListById)
	// Tasks
	web.Get("/tasks", tasks_controller.GetTasks)
	web.Post("/tasks", tasks_controller.CreateTask)
	web.Get("/tasks/:id", tasks_controller.GetTaskById)
	web.Put("/tasks/:id", tasks_controller.UpdateTaskById)
	web.Patch("/tasks/:id", tasks_controller.UpdateTaskById)
	web.Delete("/tasks/:id", tasks_controller.DeleteTaskById)
	// Users
	web.Get("/users", users_controller.GetUsers)
	web.Post("/users", users_controller.CreateUser)
	web.Get("/users/:id", users_controller.GetUserById)
	web.Put("/users/:id", users_controller.UpdateUserById)
	web.Patch("/users/:id", users_controller.UpdateUserById)
	web.Delete("/users/:id", users_controller.DeleteUserById)

	web.Run()
}
