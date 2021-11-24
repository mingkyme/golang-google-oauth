package api

import (
	"fmt"

	"example.com/session"
	"github.com/gofiber/fiber/v2"
)

func Email(c *fiber.Ctx) error {
	session, err := session.Store.Get(c)
	if err != nil {
		fmt.Println("error", err)
	}
	if session.Get("email") == nil {
		return c.SendString("No email")
	}
	email := session.Get("email").(string)
	return c.SendString(email)

}
func Logout(c *fiber.Ctx) error {
	session, err := session.Store.Get(c)
	if err != nil {
		fmt.Println("error", err)
	}
	session.Delete("email")
	session.Save()
	return c.Redirect("/")
}
