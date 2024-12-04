package services

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func LoginUser(username, password string) (string, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return "", errors.New("invalid username or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", errors.New("invalid username or password")
	}
	tokenString, err := generateJWT(username)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateJWT(username string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("JWT secret is not set")
	}

	// Create a new token object, specifying signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.New("failed to sign token")
	}

	return tokenString, nil
}
