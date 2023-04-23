package handlers

import (
	"net/http"

	"github.com/Fadiouf/goprojet/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// AuthHandler struct
type AuthHandler struct {
	DB *gorm.DB
}

// Authenticate creates a new auth
func (uh *AuthHandler) Authenticate(c echo.Context) error {
	// Bind request body to auth struct
	auth := models.AuthToken{}
	if err := c.Bind(&auth); err != nil {
		return err
	}

	// Create auth
	result := uh.DB.Create(&auth)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusCreated, auth)
}
