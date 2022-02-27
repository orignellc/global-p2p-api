package database

import (
	"fmt"
	"github.com/orignellc/global-p2p-api/models"
	"github.com/orignellc/global-p2p-api/pkg/env"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		env.GetEnv("DB_USER", ""),
		env.GetEnv("DB_PASSWORD", ""),
		env.GetEnv("DB_HOST", "127.0.0.1"),
		env.GetEnv("DB_PORT", "3306"),
		env.GetEnv("DB_NAME", ""),
	)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Err(err).
			Str("Host:Port", "127.0.0.1:3306").
			Str("Username", "root").
			Str("Password", "****").
			Msg("Connecting to Database")
	}

	err = conn.AutoMigrate(
		&models.Client{},
		&models.Transaction{},
		)
	if err != nil {
		log.Panic().
			Err(err).
			Msg("Database Migration")
	}
}
