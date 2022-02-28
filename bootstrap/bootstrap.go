package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orignellc/global-p2p-api/database"
	"github.com/orignellc/global-p2p-api/pkg/env"
	"github.com/orignellc/global-p2p-api/pkg/router"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func NewApplication() *fiber.App {
	// Get pretty print of errors from zero logger across the app
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	env.SetupEnvFile()
	database.SetupDatabase()

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	//app.Use(recover.New())
	//app.Use(logger.New())
	//app.Get("/dashboard", monitor.New())
	router.InstallRouter(app)

	return app
}