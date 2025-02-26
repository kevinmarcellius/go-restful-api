package service

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/web"
)

type ProductService interface {
	CreateProduct(ctx context.Context, request web.ProductCreateRequest) (web.ProductResponse, error)
	FindAllProducts(ctx context.Context) ([]web.ProductResponse, error)
}
