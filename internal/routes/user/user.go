package userRoutes

import (
	"github.com/gofiber/fiber/v2"
	userHandler "github.com/MuhammedSami/notes-api-fiber/internal/handlers/user"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/users")

	user.Post("/", userHandler.CreateUser)

	user.Get("/", userHandler.GetUsers)

	user.Get("/:noteId", userHandler.GetUser)
}