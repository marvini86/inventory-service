package handler

import (
	"context"
	pb "github.com/marvini86/car-service-protos/proto/inventory"
	"github.com/marvini86/inventory-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// InventoryHandler represents a handler for the inventory service
type InventoryHandler struct {
	pb.UnimplementedInventoryServiceServer
	itemService service.ItemService
}

// NewInventoryHandler creates a new InventoryHandler instance
func NewInventoryHandler(itemService service.ItemService) *InventoryHandler {
	return &InventoryHandler{itemService: itemService}
}

// CheckAvailability checks the availability of an item
func (s *InventoryHandler) CheckAvailability(ctx context.Context, req *pb.ItemRequest) (*pb.ItemResponse, error) {
	log.Printf("Checking availability for code %s", req.Code)

	if req.Code == "" {
		return nil, status.Error(codes.InvalidArgument, "Code is required")
	}

	item, err := s.itemService.GetItemAvailability(ctx, req.Code)
	if err != nil {
		log.Printf("Error checking availability for code %s: %v", req.Code, err)
		return nil, err
	}

	if item.CodeIntegration == "" {
		log.Printf("Item not found for code %s", req.Code)
		return nil, status.Error(codes.NotFound, "Item not found")
	}

	res := &pb.ItemResponse{
		Code:              item.CodeIntegration,
		Name:              item.Name,
		AvailableQuantity: int32(item.AvailableQuantity),
	}

	return res, nil
}
