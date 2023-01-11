package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	noteRoutes "github.com/MuhammedSami/notes-api-fiber/internal/routes/note"
	userRoutes "github.com/MuhammedSami/notes-api-fiber/internal/routes/user"
)

func SetupRoutes(app *fiber.App){
	api := app.Group("/api", logger.New())
	
	noteRoutes.SetupNoteRoutes(api)
	userRoutes.SetupUserRoutes(api)
}