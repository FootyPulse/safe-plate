package initializers

import "safe-plate/models"

func SyncDatabse() {
	DB.AutoMigrate(&models.User{})
}
