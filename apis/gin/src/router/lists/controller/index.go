package lists_controller

import (
	"encoding/json"
	"net/http"

	lists_service "gin-starter/src/router/lists/service"

	"github.com/gin-gonic/gin"
)

func GetLists(c *gin.Context) {
	// Get Lists
	var lists = lists_service.GetLists()
	c.JSON(http.StatusOK, lists)
}

func CreateList(c *gin.Context) {
	// Parse Body
	decoder := json.NewDecoder(c.Request.Body)
	var listRequest lists_service.ListRequest
	err := decoder.Decode(&listRequest)
	if err != nil {
		panic(err)
	}
	// Create New List
	id := lists_service.CreateList(listRequest)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func GetListById(c *gin.Context) {
	id := c.Param("id")
	// Get List
	var list = lists_service.GetList(id)
	c.JSON(http.StatusOK, list)
}

func UpdateListById(c *gin.Context) {
	id := c.Param("id")
	// Parse Body
	decoder := json.NewDecoder(c.Request.Body)
	var listRequest lists_service.ListRequest
	err := decoder.Decode(&listRequest)
	if err != nil {
		panic(err)
	}
	// Update List
	var list = lists_service.UpdateList(id, listRequest)
	c.JSON(http.StatusOK, list)
}

func PatchListById(c *gin.Context) {
	id := c.Param("id")
	// Parse Body
	decoder := json.NewDecoder(c.Request.Body)
	var listRequest lists_service.ListRequest
	err := decoder.Decode(&listRequest)
	if err != nil {
		panic(err)
	}
	// Update List
	var list = lists_service.UpdateList(id, listRequest)
	c.JSON(http.StatusOK, list)
}

func DeleteListById(c *gin.Context) {
	id := c.Param("id")
	// Get List
	var list = lists_service.DeleteList(id)
	c.JSON(http.StatusOK, list)
}
