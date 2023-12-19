package userHandlers

import (
	"github.com/Harichandra-Prasath/User-Management-System-in-GO/database"
	"github.com/Harichandra-Prasath/User-Management-System-in-GO/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Register(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error",
			"message": "Invalid Input.. Review your request body",
			"data":    err})
	}
	user.ID = uuid.New()

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error",
			"message": "Cant register the user",
			"data":    err})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User successfully registered",
		"data":    user})
}

func Dashboard(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"Status":  "error",
			"message": "No user entries found",
			"data":    nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "users found", "data": users})
}
