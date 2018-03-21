package main

import (
	"log"
	"net/http"
)

var db DB

func main() {
	db = Connect()
	router := CreateRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
