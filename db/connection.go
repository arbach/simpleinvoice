package db

import (
	"fmt"
)

type DatabaseConfig struct {
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     string `envconfig:"DB_PORT" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASS" required:"true"`
}

func ConnectionString(dbConfig DatabaseConfig) string {
	return connectionString(dbConfig.Host, dbConfig)
}

func connectionString(host string, dbConfig DatabaseConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Name,
		dbConfig.Password)
}
