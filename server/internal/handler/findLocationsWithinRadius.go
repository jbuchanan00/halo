package handler

import (
	"encoding/json"
	"halo/internal/app"
	"halo/internal/repository"
	"log"
	"math"
	"net/http"
	"strconv"
)

type rangeOfCoords struct {
	maxLat  float32
	minLat  float32
	maxLong float32
	minLong float32
}

func FindLocationsWithinRadius(a *app.App, w http.ResponseWriter, r *http.Request) {
	log.Printf("Finding locations")
	var ranges rangeOfCoords
	var locations []*app.Location
	var coords app.Coordinates

	latStr := r.URL.Query().Get("lat")
	lngStr := r.URL.Query().Get("lng")
	radStr := r.URL.Query().Get("radius")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}

	radius, err := strconv.ParseFloat(radStr, 64)
	if err != nil {
		http.Error(w, "Invalid radius", http.StatusBadRequest)
		return
	}

	coords.Latitude = float32(lat)
	coords.Longitude = float32(lng)

	ranges = *getRangeOfCoords(&coords, int16(radius))

	locations = repository.GetLocationsWithinCoords(a.DB, ranges.maxLat, ranges.minLat, ranges.maxLong, ranges.minLong)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(locations)
}

func getRangeOfCoords(origin *app.Coordinates, radius int16) *rangeOfCoords {
	var ranges rangeOfCoords

	ranges.maxLat = convertForLat(origin.Latitude, radius)
	ranges.minLat = convertForLat(origin.Latitude, -1*radius)
	ranges.maxLong = convertForLong(origin.Longitude, origin.Latitude, -1*radius)
	ranges.minLong = convertForLong(origin.Longitude, origin.Latitude, radius)

	return &ranges
}

func convertForLat(lat float32, radius int16) float32 {
	return lat + (float32(radius) / 69)
}

func convertForLong(long float32, lat float32, radius int16) float32 {
	return float32(long + (float32(radius) / float32(69*math.Cos(float64(lat)))))
}
