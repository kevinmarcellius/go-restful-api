package helper

import (
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		ProductID:   product.ProductID,
		Name:        product.Name,
		Description: product.Description,
		StockQty:    product.StockQty,
		SKU:         product.SKU,
		TaxRate:     product.TaxRate,
		CategoryId:  product.CategoryID,
	}
}

func ToProductListResponse(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}
func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
