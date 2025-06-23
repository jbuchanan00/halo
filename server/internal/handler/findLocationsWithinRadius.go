package handler

import (
	"encoding/json"
	"halo/internal/app"
	"halo/internal/repository"
	"math"
	"net/http"
)

type rangeOfCoords struct {
	maxLat  float32
	minLat  float32
	maxLong float32
	minLong float32
}

func FindLocationsWithinRadius(a *app.App, w http.ResponseWriter, r *http.Request) {
	var ranges rangeOfCoords
	var locations []*app.Location

	locations = repository.GetLocationsWithinCoords(a.DB, ranges.minLat, ranges.maxLat, ranges.minLong, ranges.maxLong)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(locations)
}

func getRangeOfCoords(origin *app.Coordinates, radius int16) *rangeOfCoords {
	var ranges rangeOfCoords

	ranges.maxLat = convertForLat(origin.Lat, radius)
	ranges.minLat = convertForLat(origin.Lat, -1*radius)
	ranges.maxLong = convertForLong(origin.Lon, origin.Lat, radius)
	ranges.minLong = convertForLong(origin.Lon, origin.Lat, radius)

	return &ranges
}

func convertForLat(lat float32, radius int16) float32 {
	return lat + (float32(radius) / 69)
}

func convertForLong(long float32, lat float32, radius int16) float32 {
	return float32(long + (float32(radius) / float32(69*math.Cos(float64(lat)))))
}
