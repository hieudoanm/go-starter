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
	lists_controller "http-starter/src/router/lists/controller"
	tasks_controller "http-starter/src/router/tasks/controller"
	users_controller "http-starter/src/router/users/controller"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// Health
	router.GET("/health", health_controller.GetHealth)
	// Tasks
	router.GET("/tasks", tasks_controller.GetTasks)
	router.POST("/tasks", tasks_controller.CreateTask)
	router.GET("/tasks/:id", tasks_controller.GetTaskById)
	router.PUT("/tasks/:id", tasks_controller.UpdateTaskById)
	router.PATCH("/tasks/:id", tasks_controller.PatchTaskById)
	router.DELETE("/tasks/:id", tasks_controller.DeleteTaskById)
	// Lists
	router.GET("/lists", lists_controller.GetLists)
	router.POST("/lists", lists_controller.CreateList)
	router.GET("/lists/:id", lists_controller.GetListById)
	router.PUT("/lists/:id", lists_controller.UpdateListById)
	router.PATCH("/lists/:id", lists_controller.PatchListById)
	router.DELETE("/lists/:id", lists_controller.DeleteListById)
	// Users
	router.GET("/users", users_controller.GetUsers)
	router.POST("/users", users_controller.CreateUser)
	router.GET("/users/:id", users_controller.GetUserById)
	router.PUT("/users/:id", users_controller.UpdateUserById)
	router.PATCH("/users/:id", users_controller.PatchUserById)
	router.DELETE("/users/:id", users_controller.DeleteUserById)
	// Server
	var PORT = "8080"
	log.Printf("🚀 Server is listening on port %s", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))

}
