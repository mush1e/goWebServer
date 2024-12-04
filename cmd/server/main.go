package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mush1e/goWebServer/internal/handlers"
)

func main() {

	// request multiplexer to match URL paths to handler functions.
	router := http.NewServeMux()

	registerRoutes(router)

	corsRouter := handlers.CORSMiddleware(router)
	loggedRouter := handlers.LoggingMiddleware(corsRouter)

	// server config
	server := &http.Server{
		Addr:         ":8080",
		Handler:      loggedRouter,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("Starting server on port 8080...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

func registerRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /health", handlers.CheckHealth)
	router.HandleFunc("GET /register", handlers.GetRegisterUser)
	router.HandleFunc("POST /register", handlers.RegisterUser)
}
