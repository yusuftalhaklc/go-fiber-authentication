package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config is a function that retrieves the value of the specified key from the configuration file (.env).
// It loads the .env file and returns the value associated with the key.
// If the .env file cannot be loaded or the key is not found, it returns an empty string.
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
