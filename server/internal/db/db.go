package db

import (
	"context"
	config "halo/internal/config"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Pool(connectionString string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.GetPostgresUrl())
	if err != nil {
		log.Printf("Unable to create connection pool: %v\n", err)
	}
	return dbpool
}
