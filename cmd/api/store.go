package main

import (
	"context"
)

type Store interface {
	GetProductList(context.Context) ([]*Product, error)
}

type MemRepository struct{}

func (p *MemRepository) GetProductList(ctx context.Context) ([]*Product, error) {

	products := []*Product{}

	products = append(products, &Product{
		ID:       1,
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
