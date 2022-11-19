package health_controller

import (
	"encoding/json"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func GetHealth(ctx *beecontext.Context) {
	var status = map[string]string{"status": "healthy"}
	byteStatus, err := json.Marshal(status)
	if err != nil {
		panic(err)
	}

	ctx.Output.Header("Content-Type", "application/json")
	ctx.Output.Body(byteStatus)
}
