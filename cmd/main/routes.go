package main

import (
	"beer-api/beers/web"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func Routes(
	cbh *web.CreateBeerHandler,
) *chi.Mux {
	mux := chi.NewMux()

	// globals middleware
	mux.Use(
		middleware.Logger,    // log every http request
		middleware.Recoverer, // recover if a panic occurs
	)

	mux.Get("/beers", nil)
	mux.Get("/hello", HelloHandler)

	mux.Post("/beers", cbh.SaveBeerHandler)

	return mux
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("done-by", "nickolaz")
	res := map[string]interface{}{"message": "hello world"}
	_ = json.NewEncoder(w).Encode(res)
}
