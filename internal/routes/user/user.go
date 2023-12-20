package userRoutes

import (
	userHandlers "github.com/Harichandra-Prasath/User-Management-System-in-GO/internal/handlers/user"
	"github.com/gofiber/fiber/v2"
)

func SetUserRoutes(router fiber.Router) {
	user := router.Group("/users")
	user.Get("/dashboard", userHandlers.Dashboard)
	user.Post("/signup", userHandlers.Signup)
	user.Post("/login", userHandlers.Login)
}
