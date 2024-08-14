package router

import (
	"fiber-postgres-REST/internal/routes/carsRoutes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	carsRoutes.SetupCarsRoutes(api)
}
