package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	// auth routes
	Route{
		"Login",
		"GET",
		"/get-token",
		GetToken,
	},
	// Todo routes
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		Validate(TodoIndex),
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		Validate(TodoShow),
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		Validate(TodoCreate),
	},
	Route{
		"TodoDelete",
		"DELETE",
		"/todos/{todoId}",
		Validate(TodoDelete),
	},
	// User routes
}
