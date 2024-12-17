package api

import (
	"log"
	"net/http"

	c "github.com/esnchez/mytheresa/internal/catalog"
)

type App struct {
	Config   Config
	Products c.Service
}

type Config struct {
	Addr     string
	DbConfig DbConfig
}

type DbConfig struct {
	Addr string
}

func NewApp(cfg Config, service c.Service) *App {
	return &App{
		Config:   cfg,
		Products: service,
	}
}

func (a *App) NewMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", a.getProductsHandler)

	return mux
}

func (a *App) Start(mux *http.ServeMux) error {

	server := &http.Server{
		Addr:    a.Config.Addr,
		Handler: mux,
	}

	log.Printf("server running on port: %s", a.Config.Addr)
	return server.ListenAndServe()
}
