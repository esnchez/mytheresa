package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var pingTimeout = 5*time.Second

func New(addr string) (*sql.DB, error){
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}