package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string

	PG_HOST     string
	PG_USER     string
	PG_PASSWORD string
	PG_DB_NAME  string
	PG_PORT     string
	PG_SSL_MODE string

	API_PREFIX string
)

func InitializeEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load Environment Variables from")
	}

	PORT = os.Getenv("PORT")

	PG_HOST = os.Getenv("PG_HOST")
	PG_USER = os.Getenv("PG_USER")
	PG_PASSWORD = os.Getenv("PG_PASSWORD")
	PG_DB_NAME = os.Getenv("PG_DB_NAME")
	PG_PORT = os.Getenv("PG_PORT")
	PG_SSL_MODE = os.Getenv("PG_SSL_MODE")

	API_PREFIX = os.Getenv("API_PREFIX")
}
