package main

import (
	"net/http"

	"github.com/gorilla/handlers"
)

const jsonHeader string = "application/json;charset=UTF-8"

func jsonResponseContentType(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonHeader)

		inner.ServeHTTP(w, r)
	})
}

func ContentType(inner http.Handler) http.Handler {
	contentTypes := "application/json"

	handler := handlers.ContentTypeHandler(inner, contentTypes)
	handler = jsonResponseContentType(handler)

	return handler
}
