package repository

import (
	"context"
	"database/sql"
	"fendi/modul-02-task/model"
	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAllProduct(ctx context.Context) ([]model.Product, error) {
	query := "SELECT id, uuid, name, stock, price FROM products"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("repository.product.GetAllProduct() Query Error: ", err.Error())
			return nil, err
		}
	}
	defer rows.Close()

	products := make([]model.Product, 0)
	for rows.Next() {
		var p model.Product
		err := rows.Scan(&p.ID, &p.UUID, &p.Name, &p.Stock, &p.Price)
		if err != nil {
			fmt.Println("repository.product.GetAllProduct() Scan Error: ", err.Error())
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) CreateProduct(ctx context.Context, p model.Product) error {
	query := "INSERT INTO products (uuid, name, stock, price) VALUES ($1, $2, $3, $4)"
	_, err := r.db.ExecContext(ctx, query, p.UUID, p.Name, p.Stock, p.Price)
	if err != nil {
		fmt.Println("repository.product.CreateProduct() Exec Error: ", err.Error())
	}

	return err
}
