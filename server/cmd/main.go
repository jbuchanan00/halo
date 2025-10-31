package main

import (
	"context"
	"fmt"
	config "halo/internal/config"
	processGeoData "halo/cmd/adhoc/processGeoData"
	"halo/internal/router"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	query := `SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'location');`

	var exists bool
	dbpool, err := pgxpool.New(context.Background(), config.GetPostgresUrl())
	if err != nil {
		fmt.Print("failed to connect to db, %w", err)
		return
	}

	dbpool.QueryRow(context.Background(), query).Scan(&exists)

	if !exists {
		processGeoData.ProcessGeoData(config.GetPostgresUrl())
	}

	r := router.New()

	port := ":8080"

	log.Printf("Starting server on port %s", port)
	err1 := http.ListenAndServe(port, r)
	if err1 != nil {
		log.Fatalf("Server failed")
	}
}
