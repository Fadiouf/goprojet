package models

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// AuthToken model
type AuthToken struct {
	ID        uint      `gorm:"primaryKey"`
	Token     string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
}

// AuthPayload represents the JWT token payload
type AuthPayload struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// JWT secret
var jwtSecret = []byte("secret")

// JWT expiration time
const jwtExpirationTime = 24 * time.Hour

// Authentication middleware
func authMiddleware(c echo.Context) (uint, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return 0, fmt.Errorf("missing Authorization header")
	}
	tokenString := authHeader[len("Bearer "):]
	token, err := jwt.ParseWithClaims(tokenString, &AuthPayload{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return 0, fmt.Errorf("invalid token: %v", err)
	}
	payload, ok := token.Claims.(*AuthPayload)
	if !ok {
		return 0, fmt.Errorf("invalid token payload")
	}
	if token.Valid && time.Now().Before(time.Unix(payload.ExpiresAt, 0)) {
		return payload.UserID, nil
	}
	return 0, fmt.Errorf("expired or invalid token")
}

func createToken(userID uint) (string, error) {
	// Define expiration time of the token
	expirationTime := time.Now().Add(5 * time.Hour)

	// Create a new JWT token with userID as the payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &AuthPayload{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
