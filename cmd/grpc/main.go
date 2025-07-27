package main

import (
	"fmt"
	pb "github.com/marvini86/car-service-protos/proto/inventory"
	"github.com/marvini86/inventory-service/internal/config"
	"github.com/marvini86/inventory-service/internal/db"
	iventory "github.com/marvini86/inventory-service/internal/grpc"
	"github.com/marvini86/inventory-service/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	config.LoadEnv()
	// Open database connection
	inventoryHandler := iventory.NewInventoryHandler(service.NewItemService(db.OpenConnection()))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterInventoryServiceServer(grpcServer, inventoryHandler)

	log.Printf("Starting gRPC server on port %s", ":50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
