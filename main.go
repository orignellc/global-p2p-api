package main

import (
	"fmt"
	"github.com/orignellc/global-p2p-api/bootstrap"
	"github.com/orignellc/global-p2p-api/pkg/env"
	"github.com/rs/zerolog/log"
)

func main() {

	app := bootstrap.NewApplication()

	err := app.Listen(
		fmt.Sprintf(
			"%s:%s",
			env.GetEnv("APP_HOST", "localhost"),
			env.GetEnv("APP_PORT", "3000"),
		),
	)
	if err != nil {
		log.Err(err).
			Msg("Serving Application")
	}
}
