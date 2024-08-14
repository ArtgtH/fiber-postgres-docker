package carsRoutes

import (
	"github.com/gofiber/fiber/v2"

	"fiber-postgres-REST/internal/handlers/carsHandler"
)

func SetupCarsRoutes(router fiber.Router) {
	note := router.Group("/cars")

	note.Post("/", carsHandler.CreateCars)

	note.Get("/", carsHandler.GetCars)

	note.Get("/:carId", carsHandler.GetCar)

	note.Put("/:carId", carsHandler.UpdateCar)

	note.Delete("/:carId", carsHandler.DeleteCar)
}
