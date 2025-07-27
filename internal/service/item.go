package service

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/marvini86/inventory-service/internal/dto"
	"github.com/marvini86/inventory-service/internal/entity"
	"gorm.io/gorm"
)

// ItemService provides methods for interacting with items
type ItemService interface {
	GetItemAvailability(ctx context.Context, code string) (item dto.AvailableItemDto, err error)
}

// itemService implements ItemService
type itemService struct {
	db *gorm.DB
}

// NewItemService creates a new ItemService instance
func NewItemService(db *gorm.DB) ItemService {
	return &itemService{db: db}
}

// GetItemAvailability gets an item by code
func (s *itemService) GetItemAvailability(ctx context.Context, code string) (item dto.AvailableItemDto, err error) {
	var itemEntity entity.Item
	err = s.db.WithContext(ctx).Where("code_integration = ?", code).First(&itemEntity).Error
	if err != nil {
		return
	}
	copier.Copy(&item, &itemEntity)
	return
}
