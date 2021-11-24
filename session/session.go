package session

import "github.com/gofiber/fiber/v2/middleware/session"

var (
	Store *session.Store
)

func Setup_Session() {
	Store = session.New()
}
