package main

import (
	"fmt"
	"net/http"
	"time"
)

func getHomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the task management API - %v\n", time.Now())
}

func getAboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a basic project im using to learn go and build some basic backend stuff in go")
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", getHomeHandler)
	router.HandleFunc("GET /about", getAboutHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("Welcome, server running on port 8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}

}
