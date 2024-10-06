package main

import (
	"gofiber-consumer/internal/config"
	"gofiber-consumer/internal/rabbitmq"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Loading configuration...")
	config.LoadConfig()

	log.Println("Starting RabbitMQ consumer...")
	go rabbitmq.StartConsumer()

	log.Println("Starting Fiber server...")
	app := fiber.New()

	// Define your routes here
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":5000"))
}
