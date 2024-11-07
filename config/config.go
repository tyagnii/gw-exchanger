// Package to read configuration files/params
package config

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func BuildConnString() string {
	builder := strings.Builder{}
	builder.WriteString("postgres://")
	builder.WriteString(os.Getenv("PGUSER"))
	builder.WriteString(":")
	builder.WriteString(os.Getenv("PGPASSWORD"))
	builder.WriteString("@")
	builder.WriteString(os.Getenv("PGHOST"))
	builder.WriteString(":")
	builder.WriteString(os.Getenv("PGPORT"))
	builder.WriteString("/")
	builder.WriteString(os.Getenv("PGDATABASE"))
	builder.WriteString("?sslmode=")
	builder.WriteString(os.Getenv("PGSSLMODE"))

	return builder.String()
}

// ReadConfig reads environment variables from config.env file
func ReadConfig(cfg string) error {
	if err := godotenv.Load(cfg); err != nil {
		return err
	}

	return nil
}
