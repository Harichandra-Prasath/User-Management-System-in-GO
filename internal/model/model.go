package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid"`
	userName  string
	firstName string
	lastName  string
	email     string
	password  string
}
