package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Post("/registration", Registration)
	app.Get("/", Hello)

	app.Listen(":3000")
}
