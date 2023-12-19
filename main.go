package main

import (
	routes "github.com/Harichandra-Prasath/User-Management-System-in-GO/Routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.SetRoutes(app)
	app.Listen(":3000")

}
