package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {
	return c.SendString("You are ready to register")
}

func LoginHandler(c *fiber.Ctx) error {
	return c.SendString("You are logged in succesfully")
}

func LogoutHandler(c *fiber.Ctx) error {
	return c.SendString("You are logged out successfully")
}

func SetRoutes(app *fiber.App) {
	app.Get("/register/", RegisterHandler)
	app.Get("/login/", LoginHandler)
	app.Get("/logout/", LogoutHandler)
}
