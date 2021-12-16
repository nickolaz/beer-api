package main

import (
	"beer-api/beers/web"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func Routes(
	cbh *web.CreateBeerHandler,
	gbh *web.GetBeerHandler,
) *chi.Mux {
	mux := chi.NewMux()

	// globals middleware
	mux.Use(
		middleware.Logger,    // log every http request
		middleware.Recoverer, // recover if a panic occurs
	)

	mux.Get("/hello", HelloHandler)

	mux.Get("/swagger/*", httpSwagger.WrapHandler)

	mux.Get("/beers", gbh.GetBeersHandler)                      //GET all the beers
	mux.Get("/beers/{beerID}/boxprice", gbh.GetBoxPriceHandler) //GET price of a box of beer
	mux.Get("/beers/day", gbh.GetBeerOfTheDayHandler)           //GET the beer of the day

	mux.Post("/beer", cbh.SaveBeerHandler)               //POST create a beer
	mux.Post("/beer/{beerID}", gbh.GetDetailBeerHandler) //POST get detail of one beer

	return mux
}

// HelloHandler godoc
// @Summary      Show a hello message for a test status of API
// @Description  Show a hello message for a test status of API
// @Tags         status
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.ErrorMsg
// @Router       /hello [get]
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("done-by", "nickolaz")
	res := map[string]interface{}{"message": "hello world"}
	_ = json.NewEncoder(w).Encode(res)
}
