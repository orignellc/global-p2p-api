package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orignellc/global-p2p-api/bootstrap"
	"github.com/rs/zerolog/log"
)

func main() {

	app := bootstrap.NewApplication()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Err(err).
			Msg("Serving Application")
	}
}
