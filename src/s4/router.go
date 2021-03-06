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
		SetItemHandler,
	},
	Route{
		"GetItem",
		"GET",
		"/items/{key}",
		GetItemHandler,
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
