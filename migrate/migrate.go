package main

import (
	"log"
	"safe-plate/initializers"
	"safe-plate/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Product{})

	if err := initializers.DB.Migrator().DropColumn(&models.Product{}, "carbohydrates"); err != nil {
		log.Fatalln("Failed to drop column Carbohydrates:", err)
	}
	if err := initializers.DB.Migrator().DropColumn(&models.Product{}, "fats"); err != nil {
		log.Fatalln("Failed to drop column Fats:", err)
	}
	if err := initializers.DB.Migrator().DropColumn(&models.Product{}, "protein"); err != nil {
		log.Fatalln("Failed to drop column Protein:", err)
	}
}
