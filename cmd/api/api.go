package main

import "net/http"

type app struct {
	config config
}

type config struct {
	addr string
}

func (a *app) start(mux *http.ServeMux) error {
	
	server := &http.Server{
		Addr: a.config.addr,
		Handler: mux,
	}
	
	return server.ListenAndServe()
}

func (a *app) NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /products", a.getProducts)

	return mux;
}

func (a *app) getProducts(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("ok"))
}