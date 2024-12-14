package main

import (
	"net/http"
)

type app struct {
	config         config
	productService *ProductService
}

type config struct {
	addr string
}

func NewApp(cfg config, productService *ProductService) *app {
	return &app{
		config:         cfg,
		productService: productService,
	}
}

func (a *app) NewMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", a.getProductsHandler)

	return mux
}

func (a *app) start(mux *http.ServeMux) error {

	server := &http.Server{
		Addr:    a.config.addr,
		Handler: mux,
	}

	return server.ListenAndServe()
}
