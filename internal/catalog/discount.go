package catalog

func CreateDiscountMap() map[string]float64 {
	discountMap := make(map[string]float64)

	discountMap["boots"] = 0.3
	discountMap["000003"] = 0.15

	return discountMap
}