package models

import (
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Handle authentication
func handleAuth(c echo.Context) error {
	// Initialize Echo
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}
	db_connexion, err := gorm.Open(postgres.Open("host=localhost user=postgres password=Projet dbname=test port=5432 sslmode=prefer"), &gorm.Config{})

	var existingUser User
	if err := db_connexion.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid email or password")
	}
	if existingUser.Password != user.Password {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid email or password")
	}
	token, err := createToken(existingUser.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
