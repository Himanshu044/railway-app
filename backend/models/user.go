package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique"`
	Password  string `valid:"required,length(6|20)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
