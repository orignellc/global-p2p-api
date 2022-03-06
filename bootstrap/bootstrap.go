package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/orignellc/global-p2p-api/database"
	"github.com/orignellc/global-p2p-api/pkg/env"
	"github.com/orignellc/global-p2p-api/pkg/router"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	zeroLog "github.com/rs/zerolog/log"
	"os"
)

func NewApplication() *fiber.App {

	// Using default logger
	logFile, err := os.OpenFile("./request.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		zeroLog.Err(err).
			Str("Log Path", "./request.log").
			Msg("Logging Request")
	}

	// Get pretty print of errors from zero logger across the app
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: logFile})

	env.SetupEnvFile()
	database.SetupDatabase()

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Middleware to recover from panics
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	// Middleware to add csrf token to request
	//app.Use(csrf.New())

	// Middleware to log http requests to request.log file
	app.Use(logger.New(logger.Config{
		Output: logFile,
	}))

	// Monitor middleware for Fiber that reports server metrics
	app.Get("/metric", monitor.New())

	// App routes
	router.InstallRouter(app)

	return app
}