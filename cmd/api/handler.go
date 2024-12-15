package main

import (
	"encoding/json"
	"net/http"

	"github.com/esnchez/mytheresa/internal/store"
)

func (a *app) getProductsHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	pag := store.Pagination{
		Limit: 5,
		Offset: 0,
	}

	pag, err := pag.ParseFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//TODO: pag struct validation

	ctx := r.Context()

	products, err := a.productService.GetProducts(ctx, pag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(products)
	if err !=  nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

