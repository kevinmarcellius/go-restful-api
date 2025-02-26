package web

import (
	"github.com/google/uuid"
)

type ProductResponse struct {
	ProductID   uuid.UUID `json:"product_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StockQty    int       `json:"stock_qty"`
	CategoryId  int       `json:"category_id"`
	SKU         int       `gorm:"not null; column:sku"`
	TaxRate     float64   `gorm:"not null; column:tax_rate"`
}
