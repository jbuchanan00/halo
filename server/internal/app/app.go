package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	DB *pgxpool.Pool
}

type Coordinates struct {
	Longitude float32
	Latitude  float32
}

type Location struct {
	Id        float32 `json:"id"`
	Name      string  `json:"name"`
	State     string  `json:"state"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Ranking   byte    `json:"ranking"`
}
