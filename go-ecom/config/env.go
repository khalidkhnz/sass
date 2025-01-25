package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// InitEnv loads environment variables from .env file
func InitEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
}


func GetPort() string {
    return ":" + getEnv("PORT", "8082")
}



// getEnv retrieves an environment variable or returns a fallback value
func getEnv(key string, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
