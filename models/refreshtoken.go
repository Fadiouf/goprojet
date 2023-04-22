package models

import "time"

// RefreshToken model
type RefreshToken struct {
	ID        uint      `gorm:"primaryKey"`
	Token     string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
}
