package migration

import (
	"safe-plate/src/domain/models"
	"safe-plate/src/infra/persistence/database"

	"gorm.io/gorm"
)

func Up1() {
	createTables(database.GetDb())
}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	//User
	tables = addNewTable(database, models.User{}, tables)

	//Post
	tables = addNewTable(database, models.Post{}, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		panic("Failed to migrate DB")
	}
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}
