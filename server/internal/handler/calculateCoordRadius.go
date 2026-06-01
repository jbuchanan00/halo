package handler

import (
	"halo/internal/app"
	"halo/internal/helpers"
	"log"
	"net/http"
	"strconv"
	"encoding/json"
)

func CalculateCoordinates(a *app.App, w http.ResponseWriter, r *http.Request){
	var coords app.Coordinates
	var ranges app.RangeOfCoords
	
	log.Printf("Calculating Coordinates")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	latStr := r.URL.Query().Get("lat")
	longStr := r.URL.Query().Get("long")
	radStr := r.URL.Query().Get("radius")

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

	radius, err := strconv.Atoi(radStr)
	if err != nil {
		http.Error(w, "Invalid Radius", http.StatusBadRequest)
		return
	}

	coords.Latitude = float32(lat)
	coords.Longitude = float32(long)

	ranges = *helpers.GetRangeOfCoords(&coords, int16(radius))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ranges)
}