package models

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Role model
type Role struct {
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"uniqueIndex;not null"`
	Description string       `gorm:"not null"`
	CreatedAt   time.Time    `gorm:"not null"`
	UpdatedAt   time.Time    `gorm:"not null"`
	DeletedAt   sql.NullTime `gorm:"index"`
	Users       []User       `gorm:"many2many:user_roles;"`
	Groups      []Group      `gorm:"many2many:group_roles;"`
}

// Handle get roles
func handleGetRoles(c echo.Context) error {
	db_connexion, err := gorm.Open(postgres.Open("host=localhost user=postgres password=Projet dbname=test port=5432 sslmode=prefer"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var roles []Role
	if err := db_connexion.Find(&roles).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, roles)
}

// Handle create role
func handleCreateRole(c echo.Context) error {
	db_connexion, err := gorm.Open(postgres.Open("host=localhost user=postgres password=Projet dbname=test port=5432 sslmode=prefer"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var role Role
	if err := c.Bind(&role); err != nil {
		return err
	}
	if err := db_connexion.Create(&role).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, role)
}

// Handle update role
func handleUpdateRole(c echo.Context) error {
	db_connexion, err := gorm.Open(postgres.Open("host=localhost user=postgres password=Projet dbname=test port=5432 sslmode=prefer"), &gorm.Config{})
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid role id")
	}
	var role Role
	if err := db_connexion.First(&role, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "role not found")
	}
	if err := c.Bind(&role); err != nil {
		return err
	}
	if err := db_connexion.Save(&role).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, role)
}

// Handle delete role
func handleDeleteRole(c echo.Context) error {
	db_connexion, err := gorm.Open(postgres.Open("host=localhost user=postgres password=Projet dbname=test port=5432 sslmode=prefer"), &gorm.Config{})
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid role id")
	}
	var role Role
	if err := db_connexion.First(&role, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "role not found")
	}
	if err := db_connexion.Delete(&role).Error; err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
