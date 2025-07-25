package handler

import (
	"encoding/json"
	"halo/internal/app"
	"halo/internal/repository"
	"log"
	"net/http"
	"strconv"
)

func ResolveCoordinates(a *app.App, w http.ResponseWriter, r *http.Request) {
	log.Printf("Resolving Coordinates")
	var location *app.Location
	w.Header().Set("Access-Control-Allow-Origin", "*")

	latStr := r.URL.Query().Get("latitude")
	longStr := r.URL.Query().Get("longitude")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "Invalid Latitude", http.StatusBadRequest)
		return
	}

	long, err := strconv.ParseFloat(longStr, 64)
	if err != nil {
		http.Error(w, "Invalid Longitude", http.StatusBadRequest)
		return
	}

	coordsForRepository := &app.Coordinates{}

	coordsForRepository.Latitude = float32(lat)

	coordsForRepository.Longitude = float32(long)

	location = repository.GetLocationByCoords(a.DB, coordsForRepository)
	log.Printf("Location %s", location.Name)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(location)
}
