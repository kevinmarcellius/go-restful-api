package domain

import "github.com/google/uuid"

type Product struct {
	ProductID   uuid.UUID `gorm:"primaryKey;column:id"`
	Name        string    `gorm:"not null;unique;column:name"`
	Description string    `gorm:"not null;column:description"`
	StockQty    int       `gorm:"not null;column:stock_qty"`
	CategoryID  int       `gorm:"not null;column:category_id"` // Ensure NOT NULL
	Category    Category  `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SKU         int       `gorm:"not null;column:sku"`
	TaxRate     float64   `gorm:"not null;column:tax_rate"`
}
