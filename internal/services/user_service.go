package services

import (
	"errors"
	"strings"

	"github.com/mush1e/goWebServer/internal/database"
	"github.com/mush1e/goWebServer/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, password string) error {
	if username == "" || password == "" {
		return errors.New("username and password cannot be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "ERROR: duplicate key value violates unique constraint") {
			return errors.New("username already taken")
		}
		return result.Error
	}

	return nil
}
