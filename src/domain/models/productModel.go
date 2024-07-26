package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Nutrition struct {
	Carbohydrates int `json:"carbohydrates"`
	Fats          int `json:"fats"`
	Protein       int `json:"protein"`
}

type Product struct {
	gorm.Model

	Price      int
	Ingredient string
	Nutrition  json.RawMessage `json:"nutrition" gorm:"type:json"`
}
