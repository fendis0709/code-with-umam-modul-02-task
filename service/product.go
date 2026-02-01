package service

import (
	"context"
	"fendi/modul-02-task/model"
	"fendi/modul-02-task/repository"
	"fendi/modul-02-task/transport"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProduct(ctx context.Context) ([]transport.ProductItemResponse, error) {
	products, err := s.repo.GetAllProduct(ctx)
	if err != nil {
		return nil, err
	}

	productsResponse := transformProduct(products)

	return productsResponse, nil
}

func transformProduct(p []model.Product) []transport.ProductItemResponse {
	var productsResponse []transport.ProductItemResponse
	for _, product := range p {
		productResponse := transport.ProductItemResponse{
			ID:          product.UUID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Category:    nil, // Assuming category data is not available in model.Product
		}
		productsResponse = append(productsResponse, productResponse)
	}

	return productsResponse
}
