package handler

import (
	"encoding/json"
	"halo/internal/app"
	"halo/internal/repository"
	"log"
	"net/http"
)

type request struct {
	Location string       `json:"location"`
	BaseLoc  app.Location `json:"baseLoc"`
}

func Autofill(a *app.App, w http.ResponseWriter, r *http.Request) {

	text := r.URL.Query().Get("text")

	if len(text) < 3 {
		log.Printf("Too short of location search")
		http.Error(w, "Location had less than 2 characters", http.StatusBadRequest)
		return
	}

	res := repository.GetLocationsLikeText(a.DB, text)

	end := 5
	if len(res) < end {
		end = len(res)
	}

	cutRes := res[0:end]

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cutRes)
}
