package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var routes = Routes{
	Route{
		"SetItem",
		"PUT",
		"/items",
		SetItem,
	},
	Route{
		"GetItem",
		"GET",
		"/items/{key}",
		GetItem,
	},
}

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
