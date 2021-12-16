package web

import (
	"beer-api/beers/gateway"
	"beer-api/external/currency"
	"beer-api/internal/database"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type BeerBox struct {
	PriceTotal float64 `json:"price_total"`
}

type GetBeerHandler struct {
	gateway.BeerGetGateway
}

func NewGetBeerHandler(client *database.PgClient) *GetBeerHandler {
	return &GetBeerHandler{gateway.GetBeerGateway(client)}
}

// GetBeersHandler godoc
// @Summary      List beers
// @Description  get all beers
// @Tags         beers
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.Beer
// @Failure      500  {object}   models.ErrorMsg
// @Failure      404  {object}   models.ErrorMsg
// @Router       /beers [get]
func (b *GetBeerHandler) GetBeersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	res, err := b.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		m := map[string]interface{}{"msg": "error in get beers"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	if res == nil {
		w.WriteHeader(http.StatusNotFound)
		m := map[string]interface{}{"msg": "error in get beers , no beers"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

// GetBeerOfTheDayHandler godoc
// @Summary      the beer of the day is a random beer
// @Description  get a random beer
// @Tags         beers
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.Beer
// @Failure      500  {object}   models.ErrorMsg
// @Failure      404  {object}   models.ErrorMsg
// @Router       /beers/day [get]
func (b *GetBeerHandler) GetBeerOfTheDayHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	res, err := b.BeerOfTheDay()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		m := map[string]interface{}{"msg": "error in get the beer"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	if res == nil {
		w.WriteHeader(http.StatusNotFound)
		m := map[string]interface{}{"msg": "error in get beers , no beers"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

// GetDetailBeerHandler godoc
// @Summary      List a detail of one beer
// @Description  get detail of one beer
// @Tags         beers
// @Accept       json
// @Produce      json
// @Param        beerID    path     integer  true  "ID of the beer"
// @Success      200  {array}   models.Beer
// @Failure      400  {object}  models.ErrorMsg
// @Failure      404  {object}  models.ErrorMsg
// @Router       /beer/{beerID} [post]
func (b *GetBeerHandler) GetDetailBeerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	beerId := chi.URLParam(r, "beerID")
	id, err := strconv.ParseUint(beerId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error beerID incorrect"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	res, err := b.GetDetails(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		m := map[string]interface{}{"msg": "error in get the beer"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

// GetBoxPriceHandler godoc
// @Summary      List price of a box by params
// @Description  get the price of a box by params
// @Tags         beers
// @Accept       json
// @Produce      json
// @Param        beerID  path      integer  true  "ID of a beer"
// @Param        currency   query     string  true  "currency for a final price"
// @Param        quantity   query     integer  false  "quantity of beer in the box"
// @Success      200      {object}  BeerBox
// @Failure      400      {object}  models.ErrorMsg
// @Failure      404      {object}  models.ErrorMsg
// @Router       /beers/{beerID}/boxprice [get]
func (b *GetBeerHandler) GetBoxPriceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	beerId := chi.URLParam(r, "beerID")
	id, err := strconv.ParseUint(beerId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error beerID incorrect"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	currencyParam := r.URL.Query().Get("currency")
	if currencyParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error currency incorrect"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	quantity := r.URL.Query().Get("quantity")
	if quantity == "" {
		quantity = "6" //Default value of a box
	}
	res, err := b.GetDetails(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		m := map[string]interface{}{"msg": "error in get the beer"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	beerCurrencyNew, err := currency.ConvertCurrencyAndPrice(res, currencyParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error get price with new currency"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	quantifyFloat, err := strconv.ParseFloat(quantity, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		m := map[string]interface{}{"msg": "error quantity not valid"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	var box BeerBox
	box.PriceTotal = beerCurrencyNew.Price * quantifyFloat
	fmt.Println(box)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&box)
}
