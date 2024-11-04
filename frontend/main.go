package main

import (
	"github.com/cicio/aiassistantapp/frontend/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	// Create views engine
	ViewsEngine := html.New("./views", ".html")

	//Start new Fiber instance
	app := fiber.New(fiber.Config{
		Views: ViewsEngine,
	})

	// Static route and directory
	app.Static("/static", "./static")

	// Create a "ping" Handler
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Fiber ")
	})

	// Create a new AppHAndler
	appHandler := handlers.NewAppHandler()

	// Use the AppHandler to handle the GET request to "/"
	app.Get("/", appHandler.HandleGetIndex)

	// Start the server on port 3000
	app.Listen(":3000")
}
