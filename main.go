package main

import (
	"log"

	"github.com/Harichandra-Prasath/User-Management-System-in-GO/database"
	"github.com/Harichandra-Prasath/User-Management-System-in-GO/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.Connect()
	router.SetRoutes(app)
	log.Fatal(app.Listen(":3000"))

}
