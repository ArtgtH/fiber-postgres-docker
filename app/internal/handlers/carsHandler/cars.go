package carsHandler

import (
	"fiber-postgres-REST/database"
	"fiber-postgres-REST/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetCars(c *fiber.Ctx) error {
	db := database.DB
	var cars []model.Car

	db.Find(&cars)

	if len(cars) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No cars present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": cars})
}

func CreateCars(c *fiber.Ctx) error {
	db := database.DB
	car := new(model.Car)

	err := c.BodyParser(&car)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	car.ID = uuid.New()
	err = db.Create(&car).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create car", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created Car", "data": car})
}

func GetCar(c *fiber.Ctx) error {
	db := database.DB
	var car model.Car

	id := c.Params("carId")

	db.Find(&car, "id = ?", id)

	if car.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No car present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Car Found", "data": car})
}

func UpdateCar(c *fiber.Ctx) error {
	type updateCar struct {
		Price     int64  `json:"price"`
		ModelName string `json:"model_name"`
	}
	db := database.DB
	var car model.Car

	id := c.Params("carId")

	db.Find(&car, "id = ?", id)

	if car.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No car present", "data": nil})
	}

	var updateCarData updateCar
	err := c.BodyParser(&updateCarData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	car.Price = updateCarData.Price
	car.ModelName = updateCarData.ModelName

	db.Save(&car)

	return c.JSON(fiber.Map{"status": "success", "message": "Car Found", "data": car})
}

func DeleteCar(c *fiber.Ctx) error {
	db := database.DB
	var car model.Car

	id := c.Params("carId")

	db.Find(&car, "id = ?", id)

	if car.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No car present", "data": nil})
	}

	err := db.Delete(&car, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete car", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Deleted car"})
}
