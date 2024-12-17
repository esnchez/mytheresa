package catalog

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Limit         int `validate:"required,gt=0,lte=5"`
	Offset        int
	Filter        string
	PriceLessThan int
}

func (p Pagination) ParseFromRequest(r *http.Request) (Pagination, error) {
	queryString := r.URL.Query()

	limit := queryString.Get("limit")
	if limit != "" {
		lim, err := strconv.Atoi(limit)
		if err != nil {
			return p, nil
		}
		p.Limit = lim
	}

	offset := queryString.Get("offset")
	if offset != "" {
		off, err := strconv.Atoi(offset)
		if err != nil {
			return p, nil
		}
		p.Offset = off
	}

	category := queryString.Get("category")
	if category != "" {
		p.Filter = category
	}

	price := queryString.Get("priceLessThan")
	if price != "" {
		pri, err := strconv.Atoi(price)
		if err != nil {
			return p, nil
		}
		p.PriceLessThan = pri
	}

	return p, nil
}