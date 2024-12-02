package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Server is up and running - %v", time.Now())
}
