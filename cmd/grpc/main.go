package main

import (
	"github.com/marvini86/inventory-service/internal/config"
	"github.com/marvini86/inventory-service/internal/server"
	"log"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	server := server.NewGrpcServer()
	if err := server.Run(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
