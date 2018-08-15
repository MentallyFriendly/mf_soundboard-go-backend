package api

import "github.com/gorilla/mux"

var (
	soundController sound
	groupController group
)

// Startup func
func Startup(r *mux.Router) {
	soundController.registerRoutes(r)
	groupController.registerRoutes(r)
}
