// Gorilla Starter:
//
//	title: Gorilla Starter
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
	health_controller "gorilla-starter/src/router/health/controller"
	lists_controller "gorilla-starter/src/router/lists/controller"
	tasks_controller "gorilla-starter/src/router/tasks/controller"
	users_controller "gorilla-starter/src/router/users/controller"

	"log"
	"time"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health", health_controller.GetHealth).Methods("GET")
	// Tasks
	router.HandleFunc("/tasks", tasks_controller.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", tasks_controller.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", tasks_controller.GetTaskById).Methods("GET")
	router.HandleFunc("/tasks/{id}", tasks_controller.UpdateTaskById).Methods("PUT")
	router.HandleFunc("/tasks/{id}", tasks_controller.PatchTaskById).Methods("PATCH")
	router.HandleFunc("/tasks/{id}", tasks_controller.DeleteTaskById).Methods("DELETE")
	// Lists
	router.HandleFunc("/lists", lists_controller.GetLists).Methods("GET")
	router.HandleFunc("/lists", lists_controller.CreateList).Methods("POST")
	router.HandleFunc("/lists/{id}", lists_controller.GetListById).Methods("GET")
	router.HandleFunc("/lists/{id}", lists_controller.UpdateListById).Methods("PUT")
	router.HandleFunc("/lists/{id}", lists_controller.PatchListById).Methods("PATCH")
	router.HandleFunc("/lists/{id}", lists_controller.DeleteListById).Methods("DELETE")
	// Users
	router.HandleFunc("/users", users_controller.GetUsers).Methods("GET")
	router.HandleFunc("/users", users_controller.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", users_controller.GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", users_controller.UpdateUserById).Methods("PUT")
	router.HandleFunc("/users/{id}", users_controller.PatchUserById).Methods("PATCH")
	router.HandleFunc("/users/{id}", users_controller.DeleteUserById).Methods("DELETE")
	// Server
	var PORT = "8080"
	var TIMEOUT = 15 * time.Second
	server := &http.Server{
		Handler:      router,
		WriteTimeout: TIMEOUT,
		ReadTimeout:  TIMEOUT,
		Addr:         "127.0.0.1:" + PORT,
	}
	log.Printf("🚀 Server is listening on port %s", PORT)
	log.Fatal(server.ListenAndServe())
}
