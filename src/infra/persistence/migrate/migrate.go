package migration

import (
	"safe-plate/src/domain/models"
	"safe-plate/src/infra/persistence/database"

	"gorm.io/gorm"
)

func Up1() {
	createTables(database.GetDb())
}

func createTables(db *gorm.DB) {
	tables := []interface{}{
		models.User{},
		models.Post{},
	}

	for _, table := range tables {
		if !db.Migrator().HasTable(table) {
			err := db.Migrator().CreateTable(table)
			if err != nil {
				panic("Failed to create table: " + err.Error())
			}
		}
	}
}
