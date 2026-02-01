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
	query := "SELECT id, uuid, name, stock, price FROM products WHERE deleted_at IS NULL ORDER BY id ASC"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []model.Product{}, nil
		}
		fmt.Println("repository.product.GetAllProduct() Query Error: ", err.Error())
		return nil, err
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

func (r *ProductRepository) GetProductByUUID(ctx context.Context, uuid string) (model.Product, error) {
	query := "SELECT id, uuid, name, stock, price FROM products WHERE uuid = $1 AND deleted_at IS NULL"
	row := r.db.QueryRowContext(ctx, query, uuid)

	var p model.Product
	err := row.Scan(&p.ID, &p.UUID, &p.Name, &p.Stock, &p.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Product{}, nil
		}
		fmt.Println("repository.product.GetProductByUUID() Scan Error: ", err.Error())
		return model.Product{}, err
	}

	return p, nil
}

func (r *ProductRepository) CreateProduct(ctx context.Context, p model.Product) error {
	query := "INSERT INTO products (uuid, name, stock, price) VALUES ($1, $2, $3, $4)"
	_, err := r.db.ExecContext(ctx, query, p.UUID, p.Name, p.Stock, p.Price)
	if err != nil {
		fmt.Println("repository.product.CreateProduct() Exec Error: ", err.Error())
	}

	return err
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, p model.Product) error {
	query := "UPDATE products SET name = $1, stock = $2, price = $3, updated_at = NOW() WHERE uuid = $4"
	_, err := r.db.ExecContext(ctx, query, p.Name, p.Stock, p.Price, p.UUID)
	if err != nil {
		fmt.Println("repository.product.UpdateProduct() Exec Error: ", err.Error())
	}

	return err
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, uuid string) error {
	query := "UPDATE products SET deleted_at = NOW() WHERE uuid = $1"
	_, err := r.db.ExecContext(ctx, query, uuid)
	if err != nil {
		fmt.Println("repository.product.DeleteProduct() Exec Error: ", err.Error())
	}

	return err
}
