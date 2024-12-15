package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/esnchez/mytheresa/internal/store"
)

type Store interface {
	GetProductList(context.Context, store.Pagination) ([]*Product, error)
}

type PostgresRepository struct {
	db *sql.DB
}

//TODO: implement final SQL query
func (ps *PostgresRepository) GetProductList(ctx context.Context, pag store.Pagination) ([]*Product, error) {

	query, args := GetQueryWithFilter(pag)

	rows, err := ps.db.QueryContext(ctx, query, args...)
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

func GetQueryWithFilter(pag store.Pagination) (string, []any) {
	args := []interface{}{}
	argIndex := 1

	query := `
		SELECT id, sku, product_name, category, price FROM products
	`
	if pag.Filter != ""{
		query += fmt.Sprintf(" WHERE category = $%d", argIndex)
		args = append(args, pag.Filter)
		argIndex++
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, pag.Limit, pag.Offset)

	return query, args
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
