package catalog

import (
	"context"
)

type MockProductService struct {
	MockGetProducts func(ctx context.Context, pag Pagination) ([]*DiscountedProduct, error)
}

func NewMockProductService() Service {
	return &MockProductService{}
}

func (m *MockProductService) GetProducts(ctx context.Context, pag Pagination) ([]*DiscountedProduct, error) {
	return m.MockGetProducts(ctx, pag)
}