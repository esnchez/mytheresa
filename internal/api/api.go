package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	shutdownCh := make(chan error)

	go func() {

		quit := make(chan os.Signal, 1)
		
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)	
		s :=<- quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		log.Println("terminate signal: ", s.String())
		
		shutdownCh <- server.Shutdown(ctx)
	}()


	log.Printf("server running on port: %s", a.Config.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	err := <- shutdownCh
	if err != nil {
		return err
	}

	return nil
}
