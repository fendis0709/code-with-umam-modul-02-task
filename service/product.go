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
		fmt.Print("s.repo.GetAllProduct() Error: ", err.Error())
		return nil, err
	}
	if len(products) == 0 {
		return []transport.ProductItemResponse{}, nil
	}

	productsResponse := transformProduct(products)

	return productsResponse, nil
}

func (s *ProductService) GetProductByUUID(ctx context.Context, uuid string) (transport.ProductItemResponse, error) {
	product, err := s.repo.GetProductByUUID(ctx, uuid)
	if err != nil {
		fmt.Print("s.repo.GetProductByUUID() Error: ", err.Error())
		return transport.ProductItemResponse{}, err
	}

	productResponse := transport.ProductItemResponse{
		ID:       product.UUID,
		Name:     product.Name,
		Stock:    product.Stock,
		Price:    product.Price,
		Category: nil, // Assuming category data is not available in model.Product
	}

	return productResponse, nil
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
	randomUUID := uuid.New().String()

	newProduct := model.Product{
		UUID:  randomUUID,
		Name:  req.Name,
		Stock: req.Stock,
		Price: req.Price,
	}

	err := s.repo.CreateProduct(ctx, newProduct)
	if err != nil {
		fmt.Print("s.repo.CreateProduct() Error: ", err.Error())
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

func (s *ProductService) UpdateProduct(ctx context.Context, id string, req transport.ProductRequest) (transport.ProductItemResponse, error) {
	newProduct := model.Product{
		UUID:  id,
		Name:  req.Name,
		Stock: req.Stock,
		Price: req.Price,
	}

	err := s.repo.UpdateProduct(ctx, newProduct)
	if err != nil {
		fmt.Print("s.repo.UpdateProduct() Error: ", err.Error())
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

func (s *ProductService) DeleteProduct(ctx context.Context, id string) error {
	err := s.repo.DeleteProduct(ctx, id)
	if err != nil {
		fmt.Print("s.repo.DeleteProduct() Error: ", err.Error())
		return err
	}

	return nil
}
