package main

import (

	"github.com/MuhammedSami/notes-api-fiber/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New();
	
	// Connect to the Database
	database.ConnectDB()
	// Send a string back for GET calls to the endpoint "/"
    app.Get("/", func(c *fiber.Ctx) error {
        err := c.SendString("And the API is UP!")
        return err
    })
	app.Listen(":3000")
	
}