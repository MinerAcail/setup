package database

import (
	"log"

	"github.com/mineracail/guardApi/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {

	// Migrate the schema
	err := db.AutoMigrate(
		
		
	)
	if err != nil {
		log.Fatal("Error migrating schema:", err)
	}

}
