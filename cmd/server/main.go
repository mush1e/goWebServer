package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mush1e/goWebServer/internal/database"
	"github.com/mush1e/goWebServer/internal/handlers"
	"github.com/mush1e/goWebServer/internal/models"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := database.Connect(); err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	if os.Getenv("ENV") == "development" {
		if err := database.DB.AutoMigrate(&models.User{}); err != nil {
			log.Fatalf("Error migrating database: %s", err)
		}
	}

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
	router.HandleFunc("GET /login", handlers.GetLoginUser)
	router.HandleFunc("POST /login", handlers.LoginUser)
}
