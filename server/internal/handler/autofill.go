package handler

import (
	"encoding/json"
	"halo/internal/app"
	"halo/internal/repository"
	"log"
	// "math"
	"net/http"
	// "slices"
)

type request struct {
	Location string       `json:"location"`
	BaseLoc  app.Location `json:"baseLoc"`
}

func Autofill(a *app.App, w http.ResponseWriter, r *http.Request) {
	var request request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Request has malformed request body %s", err)
		http.Error(w, "Body of request has an issue", http.StatusBadRequest)
		return
	}

	if len(request.Location) < 2 {
		log.Printf("Too short of location search")
		http.Error(w, "Location had less than 2 characters", http.StatusBadRequest)
		return
	}

	
	w.Header().Set("Access-Control-Allow-Origin", "*")

	res := repository.GetLocationsLikeText(a.DB, request.Location)

	end := 5
	if len(res) < end {
		end = len(res)
	}

	cutRes := res[0:end]

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cutRes)
}

// func sortLocationsByRanking(locations []*app.Location) []*app.Location {
// 	comparison := func(first *app.Location.Ranking, second *app.Location.Ranking) int {
// 		switch{
// 			case first > second:
// 				return 1
// 			case second > first:
// 				return -1:
// 			default:
// 				return 0
// 		}
// 	}

// 	slices.SortFunc(locations, comparison)

// 	return locations
// }

// func sortLocationsByDistance(baseLoc *app.Location, locations []*app.Location) []*app.Location {
// 	distCmp := func(first *app.Location, second *app.Location) int {
// 		da := distanceFromStart(baseLoc, first)
// 		db := distanceFromStart(baseLoc, second)

// 		switch {
// 		case da > db:
// 			return -1
// 		case da < db:
// 			return 1
// 		default:
// 			return 0
// 		}
// 	}

// 	slices.SortFunc(locations, distCmp)

// 	return locations
// }

// func distanceFromStart(base *app.Location, operand *app.Location) float64 {
// 	var x1, y1 float32
// 	var x2, y2 float32

// 	x1, y1 = base.Longitude, base.Latitude
// 	x2, y2 = operand.Longitude, operand.Latitude

// 	dx := float64(x2 - x1)
// 	dy := float64(y2 - y1)

// 	distance := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))

// 	return distance
// }
