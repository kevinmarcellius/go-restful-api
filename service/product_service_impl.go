package service

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/helper"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Validator         *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{productRepository, validate}
}

func (service *ProductServiceImpl) CreateProduct(ctx context.Context, request web.ProductCreateRequest) (web.ProductResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return web.ProductResponse{}, err
	}
	product := domain.Product{
		ProductID:   uuid.New(),
		Name:        request.Name,
		Description: request.Description,
		StockQty:    request.StockQty,
		SKU:         request.SKU,
		TaxRate:     request.TaxRate,
	}
	newProduct, err := service.ProductRepository.Create(ctx, product, request.CategoryId)

	if err != nil {
		return web.ProductResponse{}, err
	}
	return helper.ToProductResponse(newProduct), nil
}

func (service *ProductServiceImpl) FindAllProducts(ctx context.Context) ([]web.ProductResponse, error) {
	products, err := service.ProductRepository.FindAll(ctx)
	if err != nil {
		return []web.ProductResponse{}, err
	}

	return helper.ToProductListResponse(products), nil
}
