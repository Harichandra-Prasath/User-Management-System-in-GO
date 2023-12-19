package userHandlers

import (
	"github.com/Harichandra-Prasath/User-Management-System-in-GO/database"
	"github.com/Harichandra-Prasath/User-Management-System-in-GO/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	db := database.DB
	payload := new(model.Register)

	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error",
			"message": "Invalid Input.. Review your request body",
			"data":    err})
	}
	errors := model.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error",
			"message": "Invalid Input.. Review your request body",
			"data":    errors})
	}
	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})

	}
	hashbytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"staus": "error", "message": "Internal Server error"})
	}
	payload.Password = string(hashbytes)
	newuser := model.User{
		ID:        uuid.New(),
		UserName:  payload.Username,
		Email:     payload.Email,
		Password:  string(hashbytes),
		FirstName: payload.FirstName,
	}

	err = db.Create(&newuser).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error",
			"message": "Cant register the user",
			"data":    err})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User successfully registered",
		"data":    newuser})
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
