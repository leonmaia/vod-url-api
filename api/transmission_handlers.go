package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leonmaia/vod-api/model"
)

//GetURL ...
func (t *TransmissionHandler) GetURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transmission := t.Repository.Get(vars["id"])
	js, _ := json.Marshal(transmission)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

//Create ...
func (t *TransmissionHandler) Create(w http.ResponseWriter, r *http.Request) {
	transmission := model.Transmission{}
	err := json.NewDecoder(r.Body).Decode(&transmission)
	if err != nil {
		http.Error(w, "Please send a valid request body", 400)
		return
	}
	err = t.Repository.Insert(transmission)
	if err != nil {
		http.Error(w, "Error saving in the database", 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
