package route

import (
	"example.com/api"
	"example.com/auth"
	"github.com/gofiber/fiber/v2"
)

func Setup_Route(route fiber.Router) {
	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	route.Get("/email", api.Email)
	route.Get("/logout", api.Logout)
	route.Get("/login", auth.GoogleLogin)
	route.Get("/callback", auth.GoogleCallback)

	// route.Get("/",)
}
