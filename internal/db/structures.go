package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex;type:varchar(255);not null"`
	PasswordHash string `gorm:"type:varchar(255);not null"`
}
