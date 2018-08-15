package api

import (
	"encoding/json"
	"go_apps/go_api_apps/mf_soundboard/utils"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	soundController sound
	groupController group
)

// Startup func
func Startup(r *mux.Router) {
	soundController.registerRoutes(r)
	groupController.registerRoutes(r)
}

// Respond func
func Respond(w http.ResponseWriter, result *utils.Result) {
	if result.Error != nil {
		w.WriteHeader(result.Error.StatusCode)
		w.Header().Set("Status", http.StatusText(result.Error.StatusCode))

		data, _ := json.Marshal(result.Error)
		w.Write(data)
		return
	}

	w.WriteHeader(result.Success.StatusCode)
	w.Header().Set("Status", http.StatusText(result.Success.StatusCode))

	data, _ := json.Marshal(result.Success.Data)
	w.Write(data)
}
