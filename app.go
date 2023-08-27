package main

import (
	"log"
	"rishabhdev2700/goats/auth"
	controller "rishabhdev2700/goats/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	auth.GoogleConfig()
	app.Get("/", controller.GoogleLogin)
	app.Get("/google_callback", controller.GoogleCallback)
	log.Fatal(app.Listen(":8080"))
}
