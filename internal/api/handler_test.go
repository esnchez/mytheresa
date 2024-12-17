package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	c "github.com/esnchez/mytheresa/internal/catalog"
	u "github.com/esnchez/mytheresa/internal/utils"
)

func TestGetProductsHandler(t *testing.T) {

	tests := []struct {
		name             string
		mockService      func() c.Service
		requestQuery     string
		expectedStatus   int
		expectedResponse []c.DiscountedProduct
	}{
		{
			name: "valid request without pagination",
			mockService: func() c.Service {
				return &c.MockProductService{
					MockGetProducts: func(ctx context.Context, pag u.Pagination) ([]*c.DiscountedProduct, error) {
						return []*c.DiscountedProduct{
							{SKU: "0001", Name: "Product 1"},
							{SKU: "0002", Name: "Product 2"},
						}, nil
					},
				}
			},
			requestQuery:   "",
			expectedStatus: http.StatusOK,
			expectedResponse: []c.DiscountedProduct{
				{SKU: "0001", Name: "Product 1"},
				{SKU: "0002", Name: "Product 2"},
			},
		},
		{
			name: "valid request with pagination",
			mockService: func() c.Service {
				return &c.MockProductService{
					MockGetProducts: func(ctx context.Context, pag u.Pagination) ([]*c.DiscountedProduct, error) {
						return []*c.DiscountedProduct{
							{SKU: "0001", Name: "Product 1"},
							{SKU: "0002", Name: "Product 2"},
						}, nil
					},
				}
			},
			requestQuery:   "?limit=5&offset=0",
			expectedStatus: http.StatusOK,
			expectedResponse: []c.DiscountedProduct{
				{SKU: "0001", Name: "Product 1"},
				{SKU: "0002", Name: "Product 2"},
			},
		},
		{
			name: "service error",
			mockService: func() c.Service {
				return &c.MockProductService{
					MockGetProducts: func(ctx context.Context, pag u.Pagination) ([]*c.DiscountedProduct, error) {
						return nil, errors.New("internal server error")
					},
				}
			},
			requestQuery:   "?limit=5&offset=0",
			expectedStatus: http.StatusInternalServerError,
			expectedResponse: nil,
		},
		{
			name: "invalid pagination",
			mockService: func() c.Service {
				return &c.MockProductService{}
			},
			requestQuery:   "?limit=30", 
			expectedStatus: http.StatusBadRequest,
			expectedResponse: nil,
		},
	}

	for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T){

		mockService := tt.mockService()
		testApp := &App{Products: mockService}

		req, err := http.NewRequest(http.MethodGet, "/products"+tt.requestQuery, nil)
		if err != nil{
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		testApp.getProductsHandler(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		// mux.ServeHTTP(rec, req)

		if rec.Code != tt.expectedStatus {
			t.Errorf("expected status code to be %d and got %d", tt.expectedStatus, rec.Code)
		}
		if tt.expectedResponse != nil {
			var products []c.DiscountedProduct
			if err := json.NewDecoder(res.Body).Decode(&products); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}
			if !equalProducts(products, tt.expectedResponse) {
				t.Errorf("expected response %v, got %v", tt.expectedResponse, products)
			}
		}
	})
	}
}

func equalProducts(a, b []c.DiscountedProduct) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// func newTestApp(t *testing.T) *app{
// 	t.Helper()

// 	mockService := c.NewMockProductService()

// 	return &app{
// 		products: mockService,
// 	}
// }

