package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mush1e/goWebServer/internal/services"
)

func GetRegisterUser(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/register.html")
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	if err := services.RegisterUser(username, password); err != nil {
		if err.Error() == "username already taken" {
			w.WriteHeader(http.StatusConflict)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
}

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Server is up and running - %v", time.Now())
}
