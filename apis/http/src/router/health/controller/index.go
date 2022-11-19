package health_controller

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetHealth(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"status": "healthy"})
}
