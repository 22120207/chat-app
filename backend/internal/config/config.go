package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	MongodbUri string
	JwtSecret  string
	Database   string
}

var Conf Config

var once sync.Once

func LoadEnv() {
	once.Do(func() {
		err := godotenv.Load(`E:\Project\chat-app\backend\.env`)
		if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
		}

		mongodbUri := os.Getenv("MONGODB_URI")
		jwtSecret := os.Getenv("JWT_SECRET")
		database := os.Getenv("DATABASE")

		Conf.MongodbUri = mongodbUri
		Conf.JwtSecret = jwtSecret
		Conf.Database = database
	})
}
