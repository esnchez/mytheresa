package main

import (
	"log"

	a "github.com/esnchez/mytheresa/internal/api"
	c "github.com/esnchez/mytheresa/internal/catalog"
	"github.com/esnchez/mytheresa/internal/config"
	"github.com/esnchez/mytheresa/internal/db"
)

func main() {
	
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}

	db, err := db.New(cfg.DBAddress)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	postgresRepository := c.NewPostgresRepository(db)
	discountMap := c.CreateDiscountMap()
	productService := c.NewProductService(postgresRepository, discountMap)

	app := a.NewApp(cfg, productService)
	mux := app.NewMux()
	log.Fatal(app.Start(mux))
}



