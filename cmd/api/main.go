package main

import (
	"log"

	"ehr-headless-api/internal/config"
	"ehr-headless-api/internal/database"
	"ehr-headless-api/internal/handlers"
	"ehr-headless-api/internal/repositories"
	"ehr-headless-api/internal/services"
	"ehr-headless-api/pkg/search"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load Configuration
	config.LoadConfig()

	// Connect to Database
	database.Connect()

	// Initialize application
	app := fiber.New()
	app.Use(logger.New())

	// Dependency Injection
	searchService := search.NewSearchService()
	recordRepo := repositories.NewRecordRepository()
	recordService := services.NewRecordService(recordRepo, searchService)
	recordHandler := handlers.NewRecordHandler(recordService)

	// API v1 Group
	api := app.Group("/api/v1")

	// Routes
	api.Post("/patients/:patientId/history", recordHandler.CreateRecordEntry)

	log.Fatal(app.Listen(":3000"))
}
