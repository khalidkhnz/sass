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
    return ":" + getEnv("PORT", "3000")
}


func GetBlogPostFix(fallback string) string {
    return getEnv("BLOG_SERVICE_POSTFIX",fallback)
}

func GetBlogUrl(fallback string) string {
    return getEnv("BLOG_SERVICE_URL",fallback)
}

func GetEcomPostFix(fallback string) string {
    return getEnv("ECOM_SERVICE_POSTFIX",fallback)
}

func GetEcomUrl(fallback string) string {
    return getEnv("ECOM_SERVICE_URL",fallback)
}

func GetSassPostFix(fallback string) string {
    return getEnv("SASS_SERVICE_POSTFIX",fallback)
}

func GetSassUrl(fallback string) string {
    return getEnv("SASS_SERVICE_URL",fallback)
}




// getEnv retrieves an environment variable or returns a fallback value
func getEnv(key string, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
