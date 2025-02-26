package web

type ProductCreateRequest struct {
	Name        string  `validate:"required,min=1,max=100" json:"name"`
	Description string  `validate:"required,min=1,max=100" json:"description"`
	StockQty    int     `validate:"required,min=1,max=100" json:"stock_qty"`
	CategoryId  int     `validate:"required,min=1,max=100" json:"category_id"`
	SKU         int     `validate:"required,min=1,max=100" json:"sku"`
	TaxRate     float64 `validate:"required,min=1,max=100" json:"tax_rate"`
}
