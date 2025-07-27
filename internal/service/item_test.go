package service

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marvini86/inventory-service/internal/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetItemAvailability(t *testing.T) {
	db, mock, teardown := db.NewMockDB(t)
	defer teardown()

	ctx := context.Background()

	// Prepare mock rows
	rows := sqlmock.NewRows([]string{"code_integration", "name", "available_quantity"}).
		AddRow("000000001", "Tire", 10)

	// Expect item query
	mock.ExpectQuery(`SELECT .* FROM "items"`).
		WillReturnRows(rows)

	itemService := NewItemService(db)

	item, err := itemService.GetItemAvailability(ctx, "000000001")
	assert.NoError(t, err)
	assert.NotNil(t, item)

	assert.Equal(t, "Tire", item.Name)
	assert.Equal(t, "000000001", item.CodeIntegration)
	assert.Equal(t, 10, item.AvailableQuantity)
}

func TestGetItemAvailability_DBError(t *testing.T) {
	db, mock, cleanup := db.NewMockDB(t)
	defer cleanup()

	ctx := context.Background()

	// Simulate a db error when querying items
	mock.ExpectQuery(`SELECT .* FROM "items"`).
		WillReturnError(fmt.Errorf("mocked db failure"))

	itemService := NewItemService(db)

	_, err := itemService.GetItemAvailability(ctx, "000000001")

	assert.Error(t, err)
	assert.EqualError(t, err, "mocked db failure")
}
