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
	tasks_controller "gin-starter/src/router/tasks/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", health_controller.GetHealth)

	r.GET("/tasks", tasks_controller.GetTasks)
	r.POST("/tasks", tasks_controller.CreateTask)
	r.GET("/tasks/:id", tasks_controller.GetTaskById)
	r.PUT("/tasks/:id", tasks_controller.UpdateTaskById)
	r.PATCH("/tasks/:id", tasks_controller.PatchTaskById)
	r.DELETE("/tasks/:id", tasks_controller.DeleteTaskById)

	r.Run()
}
