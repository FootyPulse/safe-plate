package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"type:string;size:20;not null;unique"`
	Email    string `gorm:"type:string;size:64;null;unique;default:null"`
	Password string `gorm:"type:string;size:64;not null"`
}
