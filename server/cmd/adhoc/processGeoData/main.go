package main

import (
	"context"
	config "halo/cmd/adhoc/infrastructure"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

//lint:ignore U1000 this is a placeholder main function
func main() {
	dbpool, err := pgxpool.New(context.Background(), config.POSTGRES_URL)
	if err != nil {
		log.Printf("Unable to create connection pool: %v\n", err)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
	}

	log.Println(greeting)
}
