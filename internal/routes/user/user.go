package userRoutes

import (
	userHandlers "github.com/Harichandra-Prasath/User-Management-System-in-GO/internal/handlers/user"
	UserMiddleware "github.com/Harichandra-Prasath/User-Management-System-in-GO/internal/middleware/user"
	"github.com/gofiber/fiber/v2"
)

func SetUserRoutes(router fiber.Router) {
	user := router.Group("/users")
	user.Get("/dashboard", UserMiddleware.Authorize, userHandlers.Dashboard)
	user.Post("/signup", userHandlers.Signup)
	user.Post("/login", userHandlers.Login)
	user.Get("/logout", UserMiddleware.Authorize, userHandlers.Logout)
}
