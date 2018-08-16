package main

import (
	"flag"
	"fmt"
	"go_apps/go_api_apps/mf_soundboard/api"
	"go_apps/go_api_apps/mf_soundboard/db"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	seed := flag.Bool("seed", false, "Include to seed the DB")
	migrate := flag.Bool("migrate", false, "Include to run migrations on the DB")
	flag.Parse()

	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	database := db.Init(*seed, *migrate)
	defer database.Close()

	api.Startup(r)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"auth_token"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowCredentials: true,
	})

	handler := cors.Handler(s)

	r.Path("/").HandlerFunc(index).Methods("GET")

	fmt.Println("listening on port 8080..")
	http.ListenAndServe("0.0.0.0:8080", handler)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "index here :)")
}
