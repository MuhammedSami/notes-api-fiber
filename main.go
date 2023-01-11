package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/MuhammedSami/notes-api-fiber/database"
	"github.com/MuhammedSami/notes-api-fiber/router"
)

func main() {
	app := fiber.New();
	
	// Connect to the Database
	database.ConnectDB()
	
	// Setup the router
	router.SetupRoutes(app)
	app.Listen(":3000")
	
}