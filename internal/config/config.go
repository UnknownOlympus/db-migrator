package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds the configuration settings for the database connection.
// It includes the host, port, username, password, and database name.
type Config struct {
	Host     string `json:"host"`     // The hostname of the database server.
	Port     string `json:"port"`     // The port number on which the database server is listening.
	Username string `json:"username"` // The username for authenticating with the database.
	Password string `json:"password"` // The password for authenticating with the database.
	Name     string `json:"db_name"`  // The name of the database to connect to.
}

// MustLoad loads the configuration from environment variables using godotenv.
// It returns a pointer to a Config struct containing the database connection details.
// If the environment variables are not set, it will return empty values for the fields.
func MustLoad() *Config {
	_ = godotenv.Load()

	return &Config{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		Username: os.Getenv("DATABASE_USERNAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Name:     os.Getenv("DATABASE_NAME"),
	}
}
