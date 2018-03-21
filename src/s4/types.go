package main

import "net/http"

type Item struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route
