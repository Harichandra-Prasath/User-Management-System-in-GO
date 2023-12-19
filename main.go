package main

import (
	"log"

	"github.com/Harichandra-Prasath/User-Management-System-in-GO/database"
	routes "github.com/Harichandra-Prasath/User-Management-System-in-GO/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.Connect()
	routes.SetRoutes(app)
	log.Fatal(app.Listen(":3000"))

}
