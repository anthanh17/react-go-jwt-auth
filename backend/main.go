package main

import (
	"github.com/anthanh17/react-go-jwt-auth/database"
	"github.com/anthanh17/react-go-jwt-auth/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
