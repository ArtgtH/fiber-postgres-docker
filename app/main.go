package main

import (
	"fiber-postgres-REST/database"
	"fiber-postgres-REST/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	webApp := fiber.New()

	database.ConnectDB()

	webApp.Use(logger.New(logger.Config{
		Format:     "${time} ${method} ${path} - ${status} - ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
	}))

	router.SetupRoutes(webApp)

	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	logrus.Fatal(webApp.Listen(":3000"))
}
