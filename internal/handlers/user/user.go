package userHandler

import (
	"fmt"

	"github.com/MuhammedSami/notes-api-fiber/database"
	"github.com/MuhammedSami/notes-api-fiber/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No users present", "data": nil})
	}

	return c.Status(404).JSON(fiber.Map{"status": "success", "message": "Users list", "data": users})
}

func GetUser(c *fiber.Ctx) error {
	db := database.DB
	var user model.User
	id := c.Params("userId")

	fmt.Println(id)
	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found!", "data": nil})
	}

	return c.Status(404).JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Wrong User input!", "data": nil})
	}

	user.ID = uuid.New()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashedPassword[:])
	err = db.Create(&user).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create User!", "data": err})
	}

	return c.Status(404).JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}