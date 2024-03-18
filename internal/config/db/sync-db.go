package db

import "github/kodesenkoffie/server/pkg/models"

func SyncDB() {

	// -------------- Migrate Schemas -------------- \\

	// Way 1
	// DB.AutoMigrate(&models.Tenant{})

	// Way 2
	DB.AutoMigrate(new(models.Blog))
	DB.AutoMigrate(new(models.Tenant))
}
