package initializers

func SyncDatabse() {
	DB.AutoMigrate(&models.User{})
}
