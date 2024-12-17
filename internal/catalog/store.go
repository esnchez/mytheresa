package catalog

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	u "github.com/esnchez/mytheresa/internal/utils"
)

var (
	QueryTimeout = time.Second * 5
)

type Store interface {
	GetProductList(context.Context, u.Pagination) ([]*Product, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (ps *PostgresRepository) GetProductList(ctx context.Context, pag u.Pagination) ([]*Product, error) {

	query, args := getQueryWithFilters(pag)

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

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

func getQueryWithFilters(pag u.Pagination) (string, []any) {
	args := []interface{}{}
	argIndex := 1
	clauses := []string{}

	query := `
		SELECT id, sku, product_name, category, price FROM products
	`
	if pag.Filter != ""{
		clauses = append(clauses, fmt.Sprintf("category = $%d", argIndex))
		args = append(args, pag.Filter)
		argIndex++
	}

	if pag.PriceLessThan != 0 {
		clauses = append(clauses, fmt.Sprintf("price <= $%d", argIndex))
		args = append(args, pag.PriceLessThan)
		argIndex++
	}

	if len(clauses) > 0 {
		query += " WHERE " + joinClauses(clauses)
	}
	
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, pag.Limit, pag.Offset)

	return query, args
}

func joinClauses(clauses []string) string {
	result := ""
	
	for i, clause := range clauses {
		if i > 0 {
			result += " AND "
		}
		result += clause
	}
	return result
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