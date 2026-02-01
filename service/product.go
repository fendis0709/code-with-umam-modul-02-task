package service

import (
	"context"
	"fendi/modul-02-task/model"
	"fendi/modul-02-task/repository"
	"fendi/modul-02-task/transport"
	"fmt"

	"github.com/google/uuid"
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
		fmt.Print("service.product.GetAllProduct() Error: ", err.Error())
		return nil, err
	}

	productsResponse := transformProduct(products)

	return productsResponse, nil
}

func transformProduct(p []model.Product) []transport.ProductItemResponse {
	var productsResponse []transport.ProductItemResponse
	for _, product := range p {
		productResponse := transport.ProductItemResponse{
			ID:       product.UUID,
			Name:     product.Name,
			Stock:    product.Stock,
			Price:    product.Price,
			Category: nil, // Assuming category data is not available in model.Product
		}
		productsResponse = append(productsResponse, productResponse)
	}

	return productsResponse
}

func (s *ProductService) CreateProduct(ctx context.Context, req transport.ProductRequest) (transport.ProductItemResponse, error) {
	randomUUID := generateRandomUUID()

	newProduct := model.Product{
		UUID:  randomUUID,
		Name:  req.Name,
		Stock: req.Stock,
		Price: req.Price,
	}

	err := s.repo.CreateProduct(ctx, newProduct)
	if err != nil {
		fmt.Print("service.product.CreateProduct() Error: ", err.Error())
		return transport.ProductItemResponse{}, err
	}

	productResponse := transport.ProductItemResponse{
		ID:       newProduct.UUID,
		Name:     newProduct.Name,
		Stock:    newProduct.Stock,
		Price:    newProduct.Price,
		Category: nil, // Assuming category data is not available in model.Product
	}

	return productResponse, nil
}

func generateRandomUUID() string {
	return uuid.New().String()
}
