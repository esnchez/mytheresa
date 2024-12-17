package api

import (
	"net/http"

	u "github.com/esnchez/mytheresa/internal/utils"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func (a *App) getProductsHandler(w http.ResponseWriter, r *http.Request) {

	pag := u.Pagination{
		Limit:  5,
		Offset: 0,
	}

	pag, err := pag.ParseFromRequest(r)
	if err != nil {
		a.badRequest(w, err.Error())
		return
	}

	if err := Validate.Struct(pag); err != nil {
		a.badRequest(w, err.Error())
		return
	}

	products, err := a.Products.GetProducts(r.Context(), pag)
	if err != nil {
		a.internalServerError(w)
		return
	}

	if err := writeJSON(w, http.StatusOK, products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}
