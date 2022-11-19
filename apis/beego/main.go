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
	tasks_controller "bee-starter/src/router/tasks/controller"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Get("/health", health_controller.GetHealth)

	// Tasks
	web.Get("/tasks", tasks_controller.GetTasks)
	web.Post("/tasks", tasks_controller.CreateTask)
	web.Get("/tasks/:id", tasks_controller.GetTaskById)
	web.Put("/tasks/:id", tasks_controller.UpdateTaskById)
	web.Patch("/tasks/:id", tasks_controller.UpdateTaskById)
	web.Delete("/tasks/:id", tasks_controller.DeleteTaskById)

	web.Run()
}
