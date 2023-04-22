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

// User model
type User struct {
	gorm.Model
	ID         uint         `gorm:"primaryKey"`
	Name       string       `gorm:"not null"`
	Email      string       `gorm:"uniqueIndex;not null"`
	Password   string       `gorm:"not null"`
	Roles      []Role       `gorm:"many2many:user_roles;"`
	Groups     []Group      `gorm:"many2many:user_groups;"`
	CreatedAt  time.Time    `gorm:"not null"`
	UpdatedAt  time.Time    `gorm:"not null"`
	DeletedAt  sql.NullTime `gorm:"index"`
	AuthTokens []AuthToken  `gorm:"foreignKey:UserID"`
}

// Handle get users
func handleGetUsers(c echo.Context) error {
	db_connexion, err := gorm.Open(postgres.Open("host=localhost user=postgres password=Projet dbname=test port=5432 sslmode=prefer"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var users []User
	if err := db_connexion.Find(&users).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

// Handle create user
func handleCreateUser(c echo.Context) error {
	db_connexion, err := gorm.Open(postgres.Open("host=localhost user=postgres password=Projet dbname=test port=5432 sslmode=prefer"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := db_connexion.Create(&user).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// Handle update user
func handleUpdateUser(c echo.Context) error {
	db_connexion, err := gorm.Open(postgres.Open("host=localhost user=postgres password=Projet dbname=test port=5432 sslmode=prefer"), &gorm.Config{})
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}
	var user User
	if err := db_connexion.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := db_connexion.Save(&user).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// Handle delete user
func handleDeleteUser(c echo.Context) error {
	db_connexion, err := gorm.Open(postgres.Open("host=localhost user=postgres password=Projet dbname=test port=5432 sslmode=prefer"), &gorm.Config{})
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}
	var user User
	if err := db_connexion.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	if err := db_connexion.Delete(&user).Error; err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
