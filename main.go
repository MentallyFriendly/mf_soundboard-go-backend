package main

import (
	"flag"
	"fmt"
	"go_apps/go_api_apps/mf_soundboard/api"
	"go_apps/go_api_apps/mf_soundboard/config"
	"go_apps/go_api_apps/mf_soundboard/db"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	config.SetVars()
}

func main() {
	seed := flag.Bool("seed", false, "Include to seed the DB")
	migrate := flag.Bool("migrate", false, "Include to run migrations on the DB")
	flag.Parse()

	r := mux.NewRouter()

	database := db.Init(*seed, *migrate)
	defer database.Close()

	api.Startup(r)

	r.Path("/").HandlerFunc(index).Methods("GET")

	fmt.Println("listening on port 8080..")
	http.ListenAndServe("0.0.0.0:8080", r)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "index here :)")
}
