package router

import (
	"context"
	"halo/internal/app"
	config "halo/internal/config"
	"halo/internal/handler"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New() http.Handler {
	mux := http.NewServeMux()

	pool, err := pgxpool.New(context.Background(), config.GetPostgresUrl())
	if err != nil {
		log.Printf("Couldn't connect to pool")
		log.Panic()
	}

	app := &app.App{DB: pool}

	mux.HandleFunc("/", handler.Root)
	mux.HandleFunc("/autofill", func(w http.ResponseWriter, r *http.Request) {
		handler.Autofill(app, w, r)
	})
	//GET /withinradius?lat=float&lng=float&rad=int
	mux.HandleFunc("/withinradius", func(w http.ResponseWriter, r *http.Request) {
		handler.FindLocationsWithinRadius(app, w, r)
	})

	mux.HandleFunc("/resolveCoordinates", func(w http.ResponseWriter, r *http.Request) {
		handler.ResolveCoordinates(app, w, r)
	})

	mux.HandleFunc("/api/health_checks/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return mux
}
