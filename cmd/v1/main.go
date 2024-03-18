package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github/kodesenkoffie/server/cmd/v1/routes"
	"github/kodesenkoffie/server/internal/config/db"
	"github/kodesenkoffie/server/internal/config/env"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func init() {
	env.InitializeEnv()
	db.Connect()
	db.SyncDB()
}

func main() {

	app := fiber.New()

	// Logger Setup

	// Create a log file
	file, err := os.OpenFile("./.logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	// app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Output:        io.MultiWriter(os.Stdout, file),
		Format:        "[${locals:requestid}] [${pid}] [${ip}]:${port} ${status} - ${method} ${path}\n",
		DisableColors: false,
	}))

	app.Static("/", "public")
	routes.SetupRoutes(app)

	database, err := db.DB.DB()

	if err != nil {
		panic("Error closing the DB")
	}

	defer database.Close()

	app.Listen(fmt.Sprintf(":%s", env.PORT))
}
