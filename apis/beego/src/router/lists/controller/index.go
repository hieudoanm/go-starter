package lists_controller

import (
	"encoding/json"
	"strings"

	lists_service "bee-starter/src/router/lists/service"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func GetLists(ctx *beecontext.Context) {
	// Get Lists
	var lists = lists_service.GetLists()

	byteLists, byteError := json.Marshal(lists)
	if byteError != nil {
		panic(byteError)
	}

	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Body(byteLists)
}

func CreateList(ctx *beecontext.Context) {
	// Parse Body
	decoder := json.NewDecoder(ctx.Request.Body)
	var listRequest lists_service.ListRequest
	err := decoder.Decode(&listRequest)
	if err != nil {
		panic(err)
	}
	// Create New List
	id := lists_service.CreateList(listRequest)

	byteList, byteError := json.Marshal(map[string]string{"id": id})
	if byteError != nil {
		panic(byteError)
	}

	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Body(byteList)
}

func GetListById(ctx *beecontext.Context) {
	var pathParams []string = strings.Split(ctx.Request.URL.Path, "/")
	var id string = pathParams[2]
	// Get List
	var list = lists_service.GetList(id)

	byteList, byteError := json.Marshal(list)
	if byteError != nil {
		panic(byteError)
	}

	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Body(byteList)
}

func UpdateListById(ctx *beecontext.Context) {
	var pathParams []string = strings.Split(ctx.Request.URL.Path, "/")
	var id string = pathParams[2]
	// Parse Body
	decoder := json.NewDecoder(ctx.Request.Body)
	var listRequest lists_service.ListRequest
	err := decoder.Decode(&listRequest)
	if err != nil {
		panic(err)
	}
	// Update List
	var list = lists_service.UpdateList(id, listRequest)
	byteList, byteError := json.Marshal(list)
	if byteError != nil {
		panic(byteError)
	}

	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Body(byteList)
}

func PatchListById(ctx *beecontext.Context) {
	var pathParams []string = strings.Split(ctx.Request.URL.Path, "/")
	var id string = pathParams[2]
	// Parse Body
	decoder := json.NewDecoder(ctx.Request.Body)
	var listRequest lists_service.ListRequest
	err := decoder.Decode(&listRequest)
	if err != nil {
		panic(err)
	}
	// Update List
	var list = lists_service.UpdateList(id, listRequest)
	byteList, byteError := json.Marshal(list)
	if byteError != nil {
		panic(byteError)
	}

	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Body(byteList)
}

func DeleteListById(ctx *beecontext.Context) {
	var pathParams []string = strings.Split(ctx.Request.URL.Path, "/")
	var id string = pathParams[2]
	// Get List
	var list = lists_service.DeleteList(id)
	byteList, byteError := json.Marshal(list)
	if byteError != nil {
		panic(byteError)
	}

	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Body(byteList)
}
