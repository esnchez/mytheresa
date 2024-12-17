package main

import (
	"log"

	c "github.com/esnchez/mytheresa/internal/catalog"
	"github.com/esnchez/mytheresa/internal/db"
	a "github.com/esnchez/mytheresa/internal/api"
)

func main() {

	//load config from env, config pkg?
	cfg := a.Config{
		Addr: ":8080",
		DbConfig: a.DbConfig{
			Addr: "postgres://admin:admin@localhost/catalog?sslmode=disable",
		},
	}

	db, err := db.New(cfg.DbConfig.Addr)
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



