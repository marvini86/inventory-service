package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() error {
	if isDevelopment() {
		log.Println("Development mode. Loading environment variables from .env file")

		err := godotenv.Load(".env")
		if err != nil {
			return fmt.Errorf("failed to load .env file: %v", err)
		}
	}
	return nil
}

func isDevelopment() bool {
	return os.Getenv("ENV") == "" || os.Getenv("ENV") == "development"
}
