package repository

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product domain.Product, categoryId int) (domain.Product, error)
	FindAll(ctx context.Context) ([]domain.Product, error)
}
