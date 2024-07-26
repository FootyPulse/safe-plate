package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model

	Content      string
	UserID       uint `gorm:"not null"`
	User         User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LockedToTeam bool
}
