package api

import (
	"encoding/json"
	"fmt"
	"go_apps/go_api_apps/mf_soundboard/db"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type group struct{}

func (g group) registerRoutes(r *mux.Router) {
	r.Path("/groups/{name}").HandlerFunc(g.getGroup).Methods("GET")
	r.Path("/groups/{name}").HandlerFunc(g.deleteGroup).Methods("DELETE")
	r.Path("/groups/{name}").HandlerFunc(g.updateGroup).Methods("PUT", "PATCH")
	r.Path("/groups").HandlerFunc(g.getGroups).Methods("GET")
	r.Path("/groups").HandlerFunc(g.createGroup).Methods("POST")
}

func (g group) getGroup(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]

	result := db.GetGroup(name)
	data, _ := json.Marshal(result)
	w.Write(data)
}

func (g group) getGroups(w http.ResponseWriter, req *http.Request) {
	result := db.GetGroups()
	data, _ := json.Marshal(result)
	w.Write(data)
}

func (g group) createGroup(w http.ResponseWriter, req *http.Request) {
	var data map[string]*string

	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading request body ", err)
	}
	req.Body.Close()

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("Error unmarshalling bytes ", err)
	}

	result := db.CreateGroup(data)
	response, _ := json.Marshal(result)
	w.Write(response)
}

func (g group) deleteGroup(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]

	result := db.DeleteGroup(name)
	response, _ := json.Marshal(result)
	w.Write(response)
}

func (g group) updateGroup(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]

	var data map[string]*string
	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading request body ", err)
	}
	req.Body.Close()

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("Error unmarshalling bytes ", err)
	}

	result := db.UpdateGroup(name, data)
	response, _ := json.Marshal(result)
	w.Write(response)
}
