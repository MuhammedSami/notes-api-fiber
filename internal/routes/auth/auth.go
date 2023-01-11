package authRoutes

import (
	"github.com/gofiber/fiber/v2"
	noteHandler "github.com/MuhammedSami/notes-api-fiber/internal/handlers/note"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/auth")

	note.Post("/login", noteHandler.CreateNotes)

	note.Post("/register", noteHandler.GetNotes)
}