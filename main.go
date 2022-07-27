//	Gin Starter:
//		version: 0.0.1
//		title: Gin Starter
//  Schemes: http, https
//  Host: localhost:8080
//  BasePath: /
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	swagger:meta
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
  
	r.GET("/", Get)
	r.POST("/", Post)
	r.PUT("/", Put)
	r.PATCH("/", Patch)
	r.DELETE("/", Delete)

	r.Run()
}

// swagger:model HelloResponse
type HelloResponse struct {
	// Hello
	Hello string `json:"hello"`
}

// swagger:model ErrorResponse
type ErrorResponse struct {
	// Message
	Message string `json:"message"`
}

//	swagger:route GET / Get
//	Get
//
//	responses:
//		200: HelloResponse
//		400: ErrorResponse
//		401: ErrorResponse
//		404: ErrorResponse
//		500: ErrorResponse
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}

//	swagger:route POST / Post
//	Post
//
//	responses:
//		200: HelloResponse
//		400: ErrorResponse
//		401: ErrorResponse
//		404: ErrorResponse
//		500: ErrorResponse
func Post(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}

//	swagger:route PUT / put
//	Put
//
//	responses:
//		200: HelloResponse
//		400: ErrorResponse
//		401: ErrorResponse
//		404: ErrorResponse
//		500: ErrorResponse
func Put(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}

//	swagger:route PATCH / Patch
//	Patch
//
//	responses:
//		200: HelloResponse
//		400: ErrorResponse
//		401: ErrorResponse
//		404: ErrorResponse
//		500: ErrorResponse
func Patch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}

//	swagger:route DELETE / Delete
//	Delete
//
//	responses:
//		200: HelloResponse
//		400: ErrorResponse
//		401: ErrorResponse
//		404: ErrorResponse
//		500: ErrorResponse
func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}
