package server

import (
	pb "github.com/marvini86/car-service-protos/proto/inventory"
	"github.com/marvini86/inventory-service/internal/db"
	"github.com/marvini86/inventory-service/internal/handler"
	"github.com/marvini86/inventory-service/internal/service"
	"google.golang.org/grpc"
	"log"
)

// ServiceRegister represents a service register
type ServiceRegister struct {
	grpcServer *grpc.Server
}

// NewServiceRegister creates a new ServiceRegister instance
func NewServiceRegister(server *grpc.Server) *ServiceRegister {
	return &ServiceRegister{grpcServer: server}
}

// Register registers the service
func (s *ServiceRegister) Register() {
	db, err := db.OpenConnection()
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}
	pb.RegisterInventoryServiceServer(s.grpcServer, handler.NewInventoryHandler(service.NewItemService(db)))
}
