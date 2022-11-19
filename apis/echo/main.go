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
	tasks_controller "echo-starter/src/router/tasks/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
}

func Post(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
}

func GetById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{"id": id})
}

func PutById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{"id": id})
}

func PatchById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{"id": id})
}

func DeleteById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{"id": id})
}

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

	e.Logger.Fatal(e.Start(":8080"))
}
