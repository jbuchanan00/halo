package router

import (
	"halo/internal/handler"
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Root)

	return mux
}
