package main

type Product struct {
	ID       int
	SKU      string
	Name     string
	Category string
	Price    Price
}

type Price struct {
	Original int
	Final    int
	Discount *string
	Currency string
}
