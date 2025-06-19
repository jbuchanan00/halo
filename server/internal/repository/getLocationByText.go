package repository

import (
	"context"
	"halo/internal/app"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetLocationByText(db *pgxpool.Pool, loc string){
	var locations []*app.Location
	var location *app.Location

	sqlQuery := "SELECT * FROM location WHERE name LIKE \"$1% or state LIKE \"$1%"

	rows, err := db.Query(context.Background(), sqlQuery, locations)
	if err != nil {
		log.Printf("There was an error in the query %s", err)
	}

	for rows.Next() {
		
	}
}