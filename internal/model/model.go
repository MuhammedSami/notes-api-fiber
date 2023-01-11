package model

import (
	"github.com/google/uuid"
    "gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`
	Title string
	SubTitle string
	Text string
}

type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`
	Username string
	Password string
	EmailConfirmed bool `gorm:"default:true"`
	Status bool `gorm:"default:true"`
}