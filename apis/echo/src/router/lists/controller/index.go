package lists_controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	lists_service "echo-starter/src/router/lists/service"

	"github.com/labstack/echo/v4"
)

func GetLists(c echo.Context) error {
	// Get Lists
	var lists = lists_service.GetLists()
	return c.JSON(http.StatusOK, lists)
}

func CreateList(c echo.Context) error {
	// Parse Body
	decoder := json.NewDecoder(c.Request().Body)
	var listRequest lists_service.ListRequest
	err := decoder.Decode(&listRequest)
	if err != nil {
		panic(err)
	}
	// Create New List
	id := lists_service.CreateList(listRequest)
	return c.JSON(http.StatusOK, map[string]string{"id": id})
}

func GetListById(c echo.Context) error {
	var pathParams []string = strings.Split(c.Path(), "/")
	var id string = pathParams[2]
	fmt.Println(id)
	// Get List
	var list = lists_service.GetList(id)
	return c.JSON(http.StatusOK, list)
}

func UpdateListById(c echo.Context) error {
	var pathParams []string = strings.Split(c.Path(), "")
	var id string = pathParams[2]
	// Parse Body
	decoder := json.NewDecoder(c.Request().Body)
	var listRequest lists_service.ListRequest
	err := decoder.Decode(&listRequest)
	if err != nil {
		panic(err)
	}
	// Update List
	var list = lists_service.UpdateList(id, listRequest)
	return c.JSON(http.StatusOK, list)
}

func PatchListById(c echo.Context) error {
	var pathParams []string = strings.Split(c.Path(), "")
	var id string = pathParams[2]
	// Parse Body
	decoder := json.NewDecoder(c.Request().Body)
	var listRequest lists_service.ListRequest
	err := decoder.Decode(&listRequest)
	if err != nil {
		panic(err)
	}
	// Update List
	var list = lists_service.UpdateList(id, listRequest)
	return c.JSON(http.StatusOK, list)
}

func DeleteListById(c echo.Context) error {
	var pathParams []string = strings.Split(c.Path(), "")
	var id string = pathParams[2]
	// Get List
	var list = lists_service.DeleteList(id)
	return c.JSON(http.StatusOK, list)
}
