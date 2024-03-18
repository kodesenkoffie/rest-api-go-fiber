package db

import (
	"fmt"

	"github/kodesenkoffie/server/internal/config/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {

	var err error

	urlString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		env.PG_HOST,
		env.PG_USER,
		env.PG_PASSWORD,
		env.PG_DB_NAME,
		env.PG_PORT,
		env.PG_SSL_MODE,
	)

	dbConnection, err := gorm.Open(postgres.Open(urlString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	/*
		// Alternative Rather than creating dbConnection Variable assigning directly to DB variable

		DB, err = gorm.Open(postgres.Open(urlString), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		})

		tx := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	*/

	tx := dbConnection.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	if tx.Error != nil {
		panic("Unable to add UUID extension")
	}

	if err != nil {
		panic("Error connecting database")
	}

	fmt.Println("DB Connection Successful")
	fmt.Println(urlString)

	DB = dbConnection
}
