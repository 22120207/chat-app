package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(`E:\Project\chat-app\backend\.env`)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	mongodbUri := os.Getenv("MONGODB_URI")

	fmt.Printf("Database URL: %s\n", mongodbUri)
}
