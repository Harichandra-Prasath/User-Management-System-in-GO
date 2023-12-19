package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid"`
	UserName  string
	FirstName string
	LastName  string
	Email     string
	Password  string
}
