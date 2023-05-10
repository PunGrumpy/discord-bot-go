package env

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadVar(key string) string {
	// Attempt to load .env file
	_ = godotenv.Load()

	// Return the environment variable
	return os.Getenv(key)
}
