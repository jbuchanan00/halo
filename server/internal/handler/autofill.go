package handler

import (
	"encoding/json"
	"halo/internal/app"
	"log"
	"net/http"
)

type request struct {
	Location  string  `json:"location"`
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
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

	
}
