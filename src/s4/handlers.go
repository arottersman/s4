package main

import (
    "net/http"
    "fmt"
)

func SetItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This will be the set endpoint handler\n")
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This will be the get endpoint handler\n")
}