package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Pattern     string
	Method      string
	Name        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		handler = JsonHeader(handler)

		router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Pattern).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{"/users", "GET", "Users", UsersIndex},
	Route{"/users", "POST", "UsersCreate", UsersCreate},
}
