package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	DB *pgxpool.Pool
}

type Coordinates struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

type Location struct {
	Id      float32     `json:"id"`
	Name    string      `json:"name"`
	State   string      `json:"state"`
	Country string      `json:"country"`
	Coord   Coordinates `json:"coord"`
}
