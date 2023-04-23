package handlers

import (
	"net/http"
	"strconv"

	"github.com/Fadiouf/goprojet/models"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserHandler struct
type UserHandler struct {
	DB *gorm.DB
}

// CreateUser creates a new user
func (uh *UserHandler) CreateUser(c echo.Context) error {
	// Bind request body to user struct
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	// Hash user password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Create user
	result := uh.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	// Remove password from response
	user.Password = ""

	return c.JSON(http.StatusCreated, user)
}

// GetUsers returns all users
func (uh *UserHandler) GetUsers(c echo.Context) error {
	// Get all users
	users := []models.User{}
	result := uh.DB.Preload("Roles").Preload("Groups").Find(&users)
	if result.Error != nil {
		return result.Error
	}

	// Remove password from response
	for i := range users {
		users[i].Password = ""
	}

	return c.JSON(http.StatusOK, users)
}

// UpdateUser updates an existing user
func (uh *UserHandler) UpdateUser(c echo.Context) error {
	// Get user ID from URL parameter
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// Get user from database
	user := models.User{}
	result := uh.DB.Preload("Roles").Preload("Groups").First(&user, userID)
	if result.Error != nil {
		return result.Error
	}

	// Bind request body to user struct
	updatedUser := models.User{}
	if err := c.Bind(&updatedUser); err != nil {
		return err
	}

	// Update user fields
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Roles = updatedUser.Roles
	user.Groups = updatedUser.Groups

	// Save changes to database
	result = uh.DB.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	// Remove password from response
	user.Password = ""

	return c.JSON(http.StatusOK, user)
}

// DeleteUser deletes an existing user
func (uh *UserHandler) DeleteUser(c echo.Context) error {
	// Get user ID from URL parameter
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// Delete user from database
	result := uh.DB.Delete(&models.User{}, userID)
	if result.Error != nil {
		return result.Error
	}

	return c.NoContent(http.StatusNoContent)
}
