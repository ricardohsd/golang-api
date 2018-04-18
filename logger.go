package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// Logger uses gorilla LogingHandler which record following Apache style logs
func Logger(inner http.Handler) http.Handler {
	logFile := os.Stdout

	return handlers.LoggingHandler(logFile, inner)
}
