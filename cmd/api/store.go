package main

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	GetProductList(context.Context) ([]*Product, error)
}

type PostgresRepository struct {
	db *sql.DB
}

//TODO: implement final SQL query
func (ps *PostgresRepository) GetProductList(ctx context.Context) ([]*Product, error) {

	query := `
		SELECT * FROM products
	`

	rows, err := ps.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query db error: %w", err)
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		p := &Product{} 
		if err := rows.Scan(&p.ID, &p.SKU, &p.Name, &p.Category, &p.Price); err != nil {
			return nil, fmt.Errorf("row scan error: %w", err)
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return products, nil
}

type MemRepository struct{}

func (p *MemRepository) GetProductList(ctx context.Context) ([]*DiscountedProduct, error) {

	products := []*DiscountedProduct{}

	products = append(products, &DiscountedProduct{
		SKU:      "0000001",
		Name:     "test",
		Category: "test_category",
		Price: Price{
			Original: 10000,
			Final:    10000,
			Currency: "EUR",
		},
	})

	return products, nil
}
