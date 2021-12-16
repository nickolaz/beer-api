package currency

import (
	"beer-api/beers/models"
	"beer-api/internal/logs"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
)

const urlCurrencyLayer = "http://api.currencylayer.com/live?access_key=993b6bfa1b8a9e17d2faf83249012a9c"

type currencyLayerResponse struct {
	Success   bool               `json:"success"`
	Terms     string             `json:"terms"`
	Privacy   string             `json:"privacy"`
	Timestamp int64              `json:"timestamp"`
	Source    string             `json:"source"`
	Quotes    map[string]float64 `json:"quotes"`
}

func ConvertCurrencyAndPrice(b *models.Beer, currencyChange string) (*models.Beer, error) {
	currencyUrl := fmt.Sprintf("%s&currencies=%s,%s", urlCurrencyLayer, b.Currency, currencyChange)
	res, err := http.Get(currencyUrl)
	if err != nil {
		logs.Log().Error("Error get currency layer ", err)
		return nil, err
	}
	defer res.Body.Close()
	var response currencyLayerResponse
	json.NewDecoder(res.Body).Decode(&response)
	actPrice := b.Price
	if len(response.Quotes) < 2 {
		logs.Log().Error("Error currency layer , currency not found")
		err = errors.New("currency not found")
		return nil, err
	}
	var values []float64
	for _, v := range response.Quotes {
		values = append(values, v)
	}
	actPrice /= values[0]
	actPrice *= values[1]
	b.Currency = currencyChange
	b.Price = math.Floor(actPrice*100) / 100
	return b, nil
}
