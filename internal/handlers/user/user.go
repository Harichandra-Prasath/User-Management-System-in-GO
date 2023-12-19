package userHandlers

import (
	"github.com/Harichandra-Prasath/User-Management-System-in-GO/database"
	"github.com/Harichandra-Prasath/User-Management-System-in-GO/internal/model"
	"github.com/gofiber/fiber/v2"
)

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
