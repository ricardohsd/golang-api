package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	stop := make(chan os.Signal)

	signal.Notify(stop, os.Interrupt)

	logger := log.New(os.Stdout, "", 0)

	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = ":8081"
	}

	router := NewRouter()

	go func() {
		logger.Printf("Listening on http://0.0.0.0%s\n", addr)

		if err := http.ListenAndServe(addr, router); err != nil {
			logger.Fatal(err)
		}
	}()

	<-stop

	logger.Println("\nShutting down the server...")
}
