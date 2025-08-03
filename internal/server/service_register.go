package server

import (
	pb "github.com/marvini86/car-service-protos/proto/inventory"
	"github.com/marvini86/inventory-service/internal/db"
	"github.com/marvini86/inventory-service/internal/handler"
	"github.com/marvini86/inventory-service/internal/service"
	"google.golang.org/grpc"
)

// ServiceRegister represents a service register
type ServiceRegister struct {
	grpcServer *grpc.Server
}

// Register registers the service
func (s *ServiceRegister) Register() {
	pb.RegisterInventoryServiceServer(s.grpcServer, handler.NewInventoryHandler(service.NewItemService(db.OpenConnection())))
}
