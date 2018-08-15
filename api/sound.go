package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go_apps/go_api_apps/mf_soundboard/db"

	"github.com/gorilla/mux"
)

type sound struct{}

func (s sound) registerRoutes(r *mux.Router) {
	r.Path("/sounds/{id:[0-9]+}").HandlerFunc(s.getSound).Methods("GET")
	r.Path("/sounds/{id:[0-9]+}").HandlerFunc(s.deleteSound).Methods("DELETE")
	r.Path("/sounds/{id:[0-9]+}").HandlerFunc(s.updateSound).Methods("PUT", "PATCH")
	r.Path("/sounds").HandlerFunc(s.getSounds).Methods("GET")
	r.Path("/sounds").HandlerFunc(s.createSound).Methods("POST")
	r.Path("/sounds/bulk-upload").HandlerFunc(s.bulkCreateSounds).Methods("POST")
	r.Path("/sounds/").HandlerFunc(s.getSounds).Methods("GET")
}

func (s sound) getSound(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	result := db.GetSound(id)
	data, _ := json.Marshal(result)
	w.Write(data)
}

func (s sound) getSounds(w http.ResponseWriter, req *http.Request) {
	result := db.GetSounds()
	data, _ := json.Marshal(result)
	w.Write(data)
}

func (s sound) createSound(w http.ResponseWriter, req *http.Request) {
	var data map[string]*string

	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading req body ", err)
	}
	req.Body.Close()
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("Error unmarshalling []byte  ", err)
	}

	result := db.CreateSound(data)
	response, _ := json.Marshal(result)
	w.Write(response)
}

func (s sound) bulkCreateSounds(w http.ResponseWriter, req *http.Request) {
	var data []map[string]*string

	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading req body ", err)
	}
	req.Body.Close()
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("Error unmarshalling []byte  ", err)
	}

	result := db.BulkCreateSounds(data)
	response, _ := json.Marshal(result)
	w.Write(response)
}

func (s sound) deleteSound(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	result := db.DeleteSound(id)
	data, _ := json.Marshal(result)
	w.Write(data)
}

func (s sound) updateSound(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data map[string]*string

	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading req body ", err)
	}
	req.Body.Close()
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("Error unmarshalling []byte  ", err)
	}

	result := db.UpdateSound(id, data)
	response, _ := json.Marshal(result)
	w.Write(response)
}
