package health_controller

import (
	"encoding/json"
	"net/http"
)

func GetHealth(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	var status = map[string]string{"status": "healthy"}
	byteStatus, err := json.Marshal(status)
	if err != nil {
		panic(err)
	}
	writer.Write(byteStatus)

}
