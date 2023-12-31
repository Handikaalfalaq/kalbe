package database

import (
	"server/models"
	postgres "server/pkg/database"
)

func RunMigration() {
	err := postgres.DB.AutoMigrate(&models.User{}, &models.Product{})

	if err != nil {
		panic("Migration Failed")
	}
}
