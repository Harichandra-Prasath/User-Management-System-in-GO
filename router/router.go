package router

import (
	userRoutes "github.com/Harichandra-Prasath/User-Management-System-in-GO/internal/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	//setting up the individual routes
	userRoutes.SetUserRoutes(api)
}
