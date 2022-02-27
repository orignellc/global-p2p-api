package env

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var Env map[string]string

func GetEnv(key, def string) string {
	if val, ok := Env[key]; ok {
		return val
	}
	return def
}

func SetupEnvFile() {
	envFile := ".env.main"
	var err error
	Env, err = godotenv.Read(envFile)
	if err != nil {
		log.Panic().
			Err(err).
			Str("Environment File", ".env.main").
			Msg("Cannot Read Env")
	}

}