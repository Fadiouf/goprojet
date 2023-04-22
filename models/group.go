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

// Group model
type Group struct {
	ID            uint         `gorm:"primaryKey"`
	Name          string       `gorm:"uniqueIndex;not null"`
	ParentGroupID uint         `gorm:"default:null"`
	ChildGroupIDs []uint       `gorm:"-"`
	CreatedAt     time.Time    `gorm:"not null"`
	UpdatedAt     time.Time    `gorm:"not null"`
	DeletedAt     sql.NullTime `gorm:"index"`
	Users         []User       `gorm:"many2many:user_groups;"`
	Roles         []Role       `gorm:"many2many:group_roles;"`
	ParentGroup   *Group       `gorm:"foreignKey:ParentGroupID"`
	ChildGroups   []*Group     `gorm:"foreignKey:ParentGroupID"`
}

// Handle get groups
func handleGetGroups(c echo.Context) error {
	db_connexion, err := gorm.Open(postgres.Open("host=localhost user=postgres password=Projet dbname=test port=5432 sslmode=prefer"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var groups []Group
	if err := db_connexion.Find(&groups).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, groups)
}

// Handle create group
func handleCreateGroup(c echo.Context) error {
	db_connexion, err := gorm.Open(postgres.Open("host=localhost user=postgres password=Projet dbname=test port=5432 sslmode=prefer"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var group Group
	if err := c.Bind(&group); err != nil {
		return err
	}
	if err := db_connexion.Create(&group).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, group)
}

// Handle update group
func handleUpdateGroup(c echo.Context) error {
	db_connexion, err := gorm.Open(postgres.Open("host=localhost user=postgres password=Projet dbname=test port=5432 sslmode=prefer"), &gorm.Config{})
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid group id")
	}
	var group Group
	if err := db_connexion.First(&group, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "group not found")
	}
	if err := c.Bind(&group); err != nil {
		return err
	}
	if err := db_connexion.Save(&group).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, group)
}

// Handle delete group
func handleDeleteGroup(c echo.Context) error {
	db_connexion, err := gorm.Open(postgres.Open("host=localhost user=postgres password=Projet dbname=test port=5432 sslmode=prefer"), &gorm.Config{})
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid group id")
	}
	var group Group
	if err := db_connexion.First(&group, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "group not found")
	}
	if err := db_connexion.Delete(&group).Error; err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)

}
