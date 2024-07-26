package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model

	Content      string
	User         User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserID       int
	LockedToTeam bool
}
