package main

import (
	"log"

	"example.com/auth"
	"example.com/route"
	"example.com/server"
	"example.com/session"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	session.Setup_Session()

	app := server.Setup_Fiber()
	auth.Setup_GoogleOauth(app)
	route.Setup_Route(app)
	app.Listen(":80")
}
