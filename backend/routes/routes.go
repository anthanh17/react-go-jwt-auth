package routes

import (
	"github.com/anthanh17/react-go-jwt-auth/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
}
