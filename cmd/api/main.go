package main

import (
	"log"

	"github.com/esnchez/mytheresa/internal/db"
)

func main() {

	cfg := config{
		addr: ":8080",
		dbConfig: dbConfig{
			addr: "postgres://admin:admin@localhost/catalog?sslmode=disable",
		},
	}

	db, err := db.New(cfg.dbConfig.addr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	postgresRepository := &PostgresRepository{
		db: db,
	}

	discountMap := discountMap()

	productService := NewProductService(postgresRepository, discountMap)
	app := NewApp(cfg, productService)

	mux := app.NewMux()

	log.Fatal(app.start(mux))
}


func discountMap() map[string]float64 {
	discountMap := make(map[string]float64)

	discountMap["boots"] = 0.3
	discountMap["000003"] = 0.15

	return discountMap
}