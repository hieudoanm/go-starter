// HTTP Starter:
//
//	title: HTTP Starter
//	version: 0.0.1
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
	"fmt"
	"log"
	"net/http"

	health_controller "http-starter/src/router/health/controller"
	tasks_controller "http-starter/src/router/tasks/controller"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/health", health_controller.GetHealth)
	// Tasks
	router.GET("/tasks", tasks_controller.GetTasks)
	router.POST("/tasks", tasks_controller.CreateTask)
	router.GET("/tasks/:id", tasks_controller.GetTaskById)
	router.PUT("/tasks/:id", tasks_controller.UpdateTaskById)
	router.PATCH("/tasks/:id", tasks_controller.PatchTaskById)
	router.DELETE("/tasks/:id", tasks_controller.DeleteTaskById)

	var PORT = "8080"
	log.Printf("🚀 Server is listening on port %s", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))

}
