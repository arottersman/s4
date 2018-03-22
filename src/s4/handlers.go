package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func SetApplicationJson(header http.Header) {
	header.Set("Content-Type", "application/json; charset=UTF-8")
}

func SetItemHandler(w http.ResponseWriter, r *http.Request) {
	var item Item

	if r.Body == nil {
		http.Error(w, "No request body", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.SetValue(item)
	if err != nil {
		http.Error(w, "Failed to set item", http.StatusInternalServerError)
		return
	}

	SetApplicationJson(w.Header())
	w.WriteHeader(http.StatusCreated)
	return
}

func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	value, err := db.GetValue(key)
	if err == KeyNotFound {
		http.Error(w, "That key does not exist", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to get item", http.StatusInternalServerError)
		return
	}

	SetApplicationJson(w.Header())
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, value)
	return
}
