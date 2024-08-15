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
		return c.Status(404).JSON(fiber.Map{"data": "No cars present"})
	}

	var res []CarResponse
	for _, car := range cars {
		res = append(res, CarResponse{
			ID:        car.ID,
			ModelName: car.ModelName,
			Price:     car.Price,
			Country:   car.Country,
			Producer:  car.Producer,
			Date:      car.Date,
		})
	}
	return c.Status(200).JSON(fiber.Map{"data": res})
}

func CreateCar(c *fiber.Ctx) error {
	db := database.DB
	carReq := new(CreateCarRequest)

	err := c.BodyParser(&carReq)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"data": "Incorrect input"})
	}

	car := model.Car{
		ModelName: carReq.ModelName,
		Price:     carReq.Price,
		Country:   carReq.Country,
		Producer:  carReq.Producer,
		Date:      carReq.Date,
	}
	car.ID = uuid.New()
	err = db.Create(&car).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"data": err})
	}

	return c.Status(201).JSON(fiber.Map{"data": CarResponse{
		ID:        car.ID,
		ModelName: car.ModelName,
		Price:     car.Price,
		Country:   car.Country,
		Producer:  car.Producer,
		Date:      car.Date,
	}})
}

func GetCar(c *fiber.Ctx) error {
	db := database.DB
	var car model.Car
	id := c.Params("carId")

	db.Find(&car, "id = ?", id)

	if car.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"data": "No car present"})
	}

	return c.Status(200).JSON(fiber.Map{"data": CarResponse{
		ID:        car.ID,
		ModelName: car.ModelName,
		Price:     car.Price,
		Country:   car.Country,
		Producer:  car.Producer,
		Date:      car.Date,
	}})
}

func UpdateCar(c *fiber.Ctx) error {
	db := database.DB
	var car model.Car
	id := c.Params("carId")

	db.Find(&car, "id = ?", id)

	if car.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"data": "No car present"})
	}

	var updateCarData UpdateCarRequest
	err := c.BodyParser(&updateCarData)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"data": "Incorrect input"})
	}

	car.Price = updateCarData.Price
	car.ModelName = updateCarData.ModelName
	db.Save(&car)

	return c.Status(200).JSON(fiber.Map{"data": CarResponse{
		ID:        car.ID,
		ModelName: car.ModelName,
		Price:     car.Price,
		Country:   car.Country,
		Producer:  car.Producer,
		Date:      car.Date,
	}})
}

func DeleteCar(c *fiber.Ctx) error {
	db := database.DB
	var car model.Car
	id := c.Params("carId")

	db.Find(&car, "id = ?", id)

	if car.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"data": "No car present"})
	}

	err := db.Delete(&car, "id = ?", id).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"data": "Failed to delete car"})
	}

	return c.Status(200).JSON(fiber.Map{"data": "Deleted car"})
}
