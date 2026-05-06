package service

import (
	"context"
	"product-service/internal/product/dto"
	"product-service/internal/product/model"
	"product-service/internal/product/repository"
	"time"
)

type ProductService interface {
	GetAllProducts(ctx context.Context) ([]model.Product, error)
	AddProducts(ctx context.Context, username string, product dto.CreateProductRequest) error
}

type productService struct {
	r repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{
		r: r,
	}
}

func (s *productService) AddProducts(ctx context.Context, username string, product dto.CreateProductRequest) error {

	p := model.Product{
		Name:      product.Name,
		Price:     product.Price,
		Stock:     product.Stock,
		Category:  product.Category,
		CreatedAt: time.Now(),
		CreatedBy: username,
	}

	return s.r.Insert(ctx, p)
}

func (s *productService) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	return s.r.FindAll(ctx)
}
