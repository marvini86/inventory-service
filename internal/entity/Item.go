package entity

type Item struct {
	ID                int
	CodeIntegration   string
	Name              string
	Description       string
	Image             string
	AvailableQuantity int
	CategoryID        int
	CreatedAt         string
	UpdatedAt         string
}

type ItemCategory struct {
	ID        int
	Name      string
	CreatedAt string
	UpdatedAt string
}
