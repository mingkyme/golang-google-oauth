package server

import "github.com/gofiber/fiber/v2"

func Setup_Fiber() *fiber.App {
	app := fiber.New()
	return app
}
