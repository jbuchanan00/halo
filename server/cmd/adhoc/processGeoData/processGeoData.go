package prcoessGeoData

import (
	"context"
	// "encoding/json"
	"fmt"
	config "halo/internal/config"
	// "io"
	"log"
	// "os"

	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GeoCity struct {
	City       string  `json:"city"`
	StateID    string  `json:"state_id,omitempty"`
	StateName  string  `json:"state_name"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Population int64   `json:"population,omitempty"`
}

func ProcessGeoData(filePath string) error {
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, config.GetPostgresUrl())
	if err != nil {
		return fmt.Errorf("unable to create connection pool: %w", err)
	}
	defer dbpool.Close()

	// file, err := os.Open(filePath)
	// if err != nil {
	// 	return fmt.Errorf("unable to open file %s: %w", filePath, err)
	// }
	// defer file.Close()

	// bytevalue, err := io.ReadAll(file)
	// if err != nil {
	// 	return fmt.Errorf("unable to read file %s: %w", filePath, err)
	// }

	// var locations []GeoCity
	// if err = json.Unmarshal(bytevalue, &locations); err != nil {
	// 	return fmt.Errorf("unable to unmarshal json file %s: %w", filePath, err)
	// }

	if err := createTables(dbpool); err != nil {
		return fmt.Errorf("createTables failed: %w", err)
	}

	// Begin a connection for CopyFrom
	// conn, err := dbpool.Acquire(ctx)
	// if err != nil {
	// 	return fmt.Errorf("failed to acquire connection: %w", err)
	// }
	// defer conn.Release()

	// // Prepare rows to copy
	// rows := make([][]any, 0, len(locations))
	// for i, item := range locations {
	// 	rows = append(rows, []any{
	// 		i + 1,
	// 		item.City,
	// 		item.StateName,
	// 		item.Lat,
	// 		item.Lng,
	// 		item.Population,
	// 	})
	// }

	// // Use pgx.CopyFrom to load data efficiently
	// copyCount, err := conn.Conn().CopyFrom(
	// 	ctx,
	// 	pgx.Identifier{"location"},
	// 	[]string{"id", "name", "state", "latitude", "longitude", "population"},
	// 	pgx.CopyFromRows(rows),
	// )
	// if err != nil {
	// 	return fmt.Errorf("copy from failed: %w", err)
	// }

	// log.Printf("Inserted %d rows via COPY", copyCount)
	log.Printf("Successfully created table")
	return nil
}

func createTables(db *pgxpool.Pool) error {
	createSQL := `
		CREATE TABLE IF NOT EXISTS location (
		city VARCHAR NOT NULL,
		city_ascii VARCHAR,
		state_id VARCHAR,
		state_name VARCHAR,
		lat DOUBLE PRECISION,
		lng DOUBLE PRECISION,
		population BIGINT,
		ranking INTEGER,
		id INTEGER PRIMARY KEY
	)`
	_, err := db.Exec(context.Background(), createSQL)
	if err != nil {
		return fmt.Errorf("error creating tables: %w", err)
	}
	return nil
}
