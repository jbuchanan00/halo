package handler

import (
	"encoding/json"
	"halo/internal/app"
	"halo/internal/repository"
	"log"
	"math"
	"net/http"
)

type request struct {
	Location string `json:"location"`
}

func Autofill(a *app.App, w http.ResponseWriter, r *http.Request) {
	var request request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Request has malformed request body")
		http.Error(w, "Body of request has an issue", http.StatusBadRequest)
		return
	}

	if len(request.Location) < 3 {
		log.Printf("Too short of location search")
		http.Error(w, "Location had less than 3 characters", http.StatusBadRequest)
		return
	}

	res := repository.GetLocationByText(a.DB, request.Location)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func sortLocationsByDistance(starting *app.Location, locations []*app.Location) []*app.Location {
	distCmp := func(first *app.Location, second *app.Location) *app.Location {

		return derived
	}
	return locations
}

func distanceFromStart(base *app.Location, operand *app.Location) float32 {
	var x1, y1 float32
	var x2, y2 float32

	x1, y1 = base.Coord.Lon, base.Coord.Lat
	x2, y2 = operand.Coord.Lon, operand.Coord.Lat

	dx := float64(x2 - x1)
	dy := float64(y2 - y1)

	distance := math.Sqrt((dx * dx) + (dy * dy))

	return float32(distance)
}
