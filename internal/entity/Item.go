package entity

import "time"

type Item struct {
	ID                int
	CodeIntegration   string
	Name              string
	Description       string
	Image             string
	AvailableQuantity int
	CategoryID        int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type ItemCategory struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
