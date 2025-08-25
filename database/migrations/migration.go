package migrations

import (
	"fmt"
	"log"

	"my-fin-go/database"
	"my-fin-go/models"
)

func Migration() {
	err := database.DB.AutoMigrate(
		&models.User{},
		&models.Finance{},
	)
	if err != nil {
		log.Fatal("Failed to migrate...")
	}

	fmt.Println("Migrated successfully")
}
