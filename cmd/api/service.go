package main

import "context"

type ProductService struct {
	store Store
}

func NewProductService(store Store) *ProductService {
	return &ProductService{
		store: store,
	}
}

func (p *ProductService) GetProducts(ctx context.Context) ([]*Product, error) {
	return p.store.GetProductList(ctx)
}
