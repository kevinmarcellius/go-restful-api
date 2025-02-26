package repository

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/domain"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (p *ProductRepositoryImpl) Create(ctx context.Context, product domain.Product, categoryId int) (domain.Product, error) {
	// get category from db
	var category domain.Category
	if err := p.db.First(&category, categoryId).Error; err != nil {
		return domain.Product{}, err
	}
	product.Category = category
	if err := p.db.WithContext(ctx).Create(&product).Error; err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (p *ProductRepositoryImpl) FindAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	if err := p.db.WithContext(ctx).Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}
