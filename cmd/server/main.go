package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Placeholder route
func checkHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Server is up and running")
}

func main() {

	// request multiplexer to match URL paths to handler functions.
	router := http.NewServeMux()

	router.HandleFunc("/health", checkHealth)

	// server config
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("Starting server on port 8080...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
