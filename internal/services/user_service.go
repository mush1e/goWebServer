package services

import (
	"errors"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username  string
	Password  string
	CreatedAt time.Time
}

var (
	users      = make(map[string]User)
	usersMutex = sync.Mutex{}
)

func RegisterUser(username, password string) error {
	if username == "" || password == "" {
		return errors.New("username and password cannot be empty")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Lock the map to ensure thread safety
	usersMutex.Lock()
	defer usersMutex.Unlock()

	// Check if the username already exists
	if _, exists := users[username]; exists {
		return errors.New("username already taken")
	}

	// Store the new user with the hashed password
	users[username] = User{
		Username:  username,
		Password:  string(hashedPassword), // Store hashed password
		CreatedAt: time.Now(),
	}

	return nil
}
