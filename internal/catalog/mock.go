package catalog

import (
	"context"
	u "github.com/esnchez/mytheresa/internal/utils"

)

type MockProductService struct {
	MockGetProducts func(ctx context.Context, pag u.Pagination) ([]*DiscountedProduct, error)
}

func NewMockProductService() Service {
	return &MockProductService{}
}

func (m *MockProductService) GetProducts(ctx context.Context, pag u.Pagination) ([]*DiscountedProduct, error) {
	return m.MockGetProducts(ctx, pag)
}