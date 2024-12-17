package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var pingTimeout = 5*time.Second

func New(addr string) (*sql.DB, error){
	var db *sql.DB
	var err error

	log.Println(addr)

	for i := 0; i < 5; i++ {

		ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
		defer cancel()

		db, err = sql.Open("postgres", addr)
		if err == nil && db.PingContext(ctx) == nil {
			log.Println("Connected to Postgres!")
			break
		}
		log.Println("database not ready, retrying in 2 seconds...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to Postgres after retries: %v", err)
	}

	return db, nil
}
	