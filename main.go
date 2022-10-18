package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/high-performance-payment-gateway/balance-service/balance"
)

func main() {
	app := fiber.New()
	// Routes
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World 👋!")
	})
	// Start server
	//app.Listen(":8080")

	balanceModule := balance.NewModule()
	balanceModule.Start()
}
