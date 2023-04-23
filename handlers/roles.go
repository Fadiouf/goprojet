package handlers

import (
	"net/http"
	"strconv"

	"github.com/Fadiouf/goprojet/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// RoleHandler struct
type RoleHandler struct {
	DB *gorm.DB
}

// CreateRole creates a new role
func (uh *RoleHandler) CreateRole(c echo.Context) error {
	// Bind request body to role struct
	role := models.Role{}
	if err := c.Bind(&role); err != nil {
		return err
	}

	// Create role
	result := uh.DB.Create(&role)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusCreated, role)
}

// GetRoles returns all roles
func (uh *RoleHandler) GetRoles(c echo.Context) error {
	// Get all roles
	roles := []models.Role{}
	result := uh.DB.Preload("Users").Preload("Groups").Find(&roles)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, roles)
}

// UpdateRole updates an existing role
func (uh *RoleHandler) UpdateRole(c echo.Context) error {
	// Get role ID from URL parameter
	roleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// Get role from database
	role := models.Role{}
	result := uh.DB.Preload("Users").Preload("Groups").First(&role, roleID)
	if result.Error != nil {
		return result.Error
	}

	// Bind request body to role struct
	updatedRole := models.Role{}
	if err := c.Bind(&updatedRole); err != nil {
		return err
	}

	// Update user fields
	role.Name = updatedRole.Name
	role.Users = updatedRole.Users
	role.Groups = updatedRole.Groups
	role.Description = updatedRole.Description

	// Save changes to database
	result = uh.DB.Save(&role)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, role)
}

// DeleteRoles deletes an existing role
func (uh *RoleHandler) DeleteRole(c echo.Context) error {
	// Get role ID from URL parameter
	roleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// Delete role from database
	result := uh.DB.Delete(&models.Role{}, roleID)
	if result.Error != nil {
		return result.Error
	}

	return c.NoContent(http.StatusNoContent)
}
