package main

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func ContentType(inner http.Handler) http.Handler {
	contentTypes := "application/json"
	return handlers.ContentTypeHandler(inner, contentTypes)
}
