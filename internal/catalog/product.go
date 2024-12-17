package catalog

type Product struct {
	ID       int
	SKU      string
	Name     string
	Category string
	Price    int
}

type DiscountedProduct struct {
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    Price  `json:"price"`
}

type Price struct {
	Original int     `json:"original"`
	Final    int     `json:"final"`
	Discount *string `json:"discount"`
	Currency string  `json:"currency"`
}
