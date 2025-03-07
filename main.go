package main

import (
	"safe-plate/src/api"
	"safe-plate/src/config"
	"safe-plate/src/infra/persistence/database"
	migration "safe-plate/src/infra/persistence/migrate"
	"safe-plate/src/initializers"
)

func main() {

	initializers.LoadEnvVariables()

	cfg := config.GetConfig()

	err := database.InitDb(cfg)

	if err != nil {
		panic("failed to load databse")
	}

	migration.Up1()

	api.InitServer()

}
