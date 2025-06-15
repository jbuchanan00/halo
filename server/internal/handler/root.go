package handler

import (
	"encoding/json"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	resp := "Hello from root"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
