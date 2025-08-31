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
}

var conf Config

var once sync.Once

func LoadEnv() Config {
	once.Do(func() {
		err := godotenv.Load(`E:\Project\chat-app\backend\.env`)
		if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
		}

		mongodbUri := os.Getenv("MONGODB_URI")
		jwtSecret := os.Getenv("JWT_SECRET")

		conf.MongodbUri = mongodbUri
		conf.JwtSecret = jwtSecret
	})

	return conf
}
