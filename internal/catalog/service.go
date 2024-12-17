package catalog

import (
	"context"
	"fmt"
	"log"
	"math"
)

var currency = "EUR"

type Service interface {
	GetProducts(ctx context.Context, pag Pagination) ([]*DiscountedProduct, error)
}

type ProductService struct {
	store      Store
	discounter map[string]float64
}

func NewProductService(store Store, discounter map[string]float64) *ProductService {
	return &ProductService{
		store:      store,
		discounter: discounter,
	}
}

func (p *ProductService) GetProducts(ctx context.Context, pag Pagination) ([]*DiscountedProduct, error) {

	products, err := p.store.GetProductList(ctx, pag)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return nil, err
	}

	discountedProducts := p.ApplyDiscounts(products)

	return discountedProducts, nil
}

func (p *ProductService) ApplyDiscounts(products []*Product) []*DiscountedProduct {
	discountedProducts := make([]*DiscountedProduct, 0, len(products))

	for _, product := range products {
		var finalDiscount float64
		var discountPercentage *string

		if discount, ok := p.discounter[product.Category]; ok {
			finalDiscount = math.Max(finalDiscount, discount)
		}

		if discount, ok := p.discounter[product.SKU]; ok {
			finalDiscount = math.Max(finalDiscount, discount)
		}

		if finalDiscount > 0 {
			percentage := fmt.Sprintf("%.0f%%", finalDiscount*100)
			discountPercentage = &percentage
		}

		discountedProduct := &DiscountedProduct{
			SKU:      product.SKU,
			Name:     product.Name,
			Category: product.Category,
			Price: Price{
				Original: product.Price,
				Final:    int(float64(product.Price) * (1 - finalDiscount)),
				Discount: discountPercentage,
				Currency: currency,
			},
		}

		discountedProducts = append(discountedProducts, discountedProduct)
	}

	return discountedProducts
}
