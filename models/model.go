package models

import (
	"time"
)

// model User
type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Email     string `gorm:"unique"`
	Password  string
	Photo     []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// ID        uint      `gorm:"primary_key"`
	// Username  string    `gorm:"not null"`
	// Email     string    `gorm:"unique;not null"`
	// Password  string    `gorm:"not null;min:6"`
	// PhotoID   uint      `gorm:"foreignkey:PhotoID;references:ID"`
	// Photo     *Photo    `gorm:"foreignkey:ID;references:PhotoID"`
	// CreatedAt time.Time `gorm:"not null"`
	// UpdatedAt time.Time `gorm:"not null"`
}

// model for photos.

type Photo struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	// Add photo fields here
	CreatedAt time.Time
	UpdatedAt time.Time
	// ID        uint      `gorm:"primary_key"`
	// Name      string    `gorm:"not null"`
	// Path      string    `gorm:"not null"`
	// UserID    uint      `gorm:"not null"`
	// User      *User     `gorm:"foreignkey:UserID;references:ID"`
	// CreatedAt time.Time `gorm:"not null"`
	// UpdatedAt time.Time `gorm:"not null"`
}

type Credentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
