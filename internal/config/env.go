package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	if isDevelopment() {
		log.Println("Development mode. Loading environment variables from .env file")

		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func isDevelopment() bool {
	return os.Getenv("ENV") == "" || os.Getenv("ENV") == "development"
}
