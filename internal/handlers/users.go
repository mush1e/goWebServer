package handlers

import (
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

// var (
// 	users      = make(map[string]User)
// 	usersMutex = sync.Mutex{}
// )

func RegisterUser(w http.ResponseWriter, r *http.Request) {

}

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Server is up and running - %v", time.Now())
}
