package main

import "log"

func main() {

	cfg := config{
		addr: ":8080",
	}

	memRepository := &MemRepository{}
	productService := NewProductService(memRepository)
	app := NewApp(cfg, productService)

	mux := app.NewMux()

	log.Fatal(app.start(mux))
}