package handlers

import (
	"net/http"
	"strconv"

	"github.com/Fadiouf/goprojet/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// GroupHandler struct
type GroupHandler struct {
	DB *gorm.DB
}

// CreateGroup creates a new group
func (uh *GroupHandler) CreateGroup(c echo.Context) error {
	// Bind request body to group struct
	group := models.Group{}
	if err := c.Bind(&group); err != nil {
		return err
	}

	// Create group
	result := uh.DB.Create(&group)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusCreated, group)
}

// GetUsers returns all groups
func (uh *GroupHandler) GetGroups(c echo.Context) error {
	// Get all groups
	groups := []models.Group{}
	result := uh.DB.Preload("Roles").Preload("Users").Find(&groups)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, groups)
}

// UpdateGroup updates an existing group
func (uh *GroupHandler) UpdateGroup(c echo.Context) error {
	// Get group ID from URL parameter
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// Get group from database
	group := models.Group{}
	result := uh.DB.Preload("Roles").Preload("Users").First(&group, groupID)
	if result.Error != nil {
		return result.Error
	}

	// Bind request body to group struct
	updatedGroup := models.Group{}
	if err := c.Bind(&updatedGroup); err != nil {
		return err
	}

	// Update group fields
	group.Name = updatedGroup.Name
	group.Roles = updatedGroup.Roles
	group.Users = updatedGroup.Users

	// Save changes to database
	result = uh.DB.Save(&group)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, group)
}

// DeleteGroup deletes an existing group
func (uh *GroupHandler) DeleteGroup(c echo.Context) error {
	// Get group ID from URL parameter
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// Delete group from database
	result := uh.DB.Delete(&models.Group{}, groupID)
	if result.Error != nil {
		return result.Error
	}

	return c.NoContent(http.StatusNoContent)
}
