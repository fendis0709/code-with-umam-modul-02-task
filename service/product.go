package service

import (
	"context"
	"fendi/modul-02-task/model"
	"fendi/modul-02-task/repository"
	"fendi/modul-02-task/transport"

	"github.com/google/uuid"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProduct(ctx context.Context) (transport.ProductsResponse, error) {
	products, err := s.repo.GetAllProduct(ctx)
	if err != nil {
		return transport.ProductsResponse{}, err
	}

	productsResponse := transformProduct(products)

	return transport.ProductsResponse{
		Code:    200,
		Message: "Products retrieved successfully",
		Data:    transport.ProductsResponseData{Products: productsResponse},
	}, nil
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

func (s *ProductService) CreateProduct(ctx context.Context, req transport.PackageRequest) (transport.ProductResponse, error) {
	randomUUID := generateRandomUUID()

	newProduct := model.Product{
		UUID:        randomUUID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	err := s.repo.CreateProduct(ctx, newProduct)
	if err != nil {
		return transport.ProductResponse{}, err
	}

	productResponse := transport.ProductItemResponse{
		ID:          newProduct.UUID,
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		Category:    nil, // Assuming category data is not available in model.Product
	}

	return transport.ProductResponse{
		Code:    201,
		Message: "Product created successfully",
		Data:    transport.ProductResponseData{Product: productResponse},
	}, nil
}

func generateRandomUUID() string {
	return uuid.New().String()
}
