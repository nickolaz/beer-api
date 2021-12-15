package web

import (
	"beer-api/beers/gateway"
	"beer-api/beers/models"
	"beer-api/internal/database"
	"encoding/json"
	"net/http"
)

type CreateBeerHandler struct {
	gateway.BeerCreateGateway
}

func NewCreateBeerHandler(client *database.PgClient) *CreateBeerHandler {
	return &CreateBeerHandler{gateway.NewBeerCreateGateway(client)}
}

func ParseRequest(r *http.Request) *models.CreateBeerCMD {
	body := r.Body
	defer body.Close()
	var beer models.CreateBeerCMD
	json.NewDecoder(body).Decode(&beer)
	return &beer
}

func (b *CreateBeerHandler) SaveBeerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	beer := ParseRequest(r)
	res, err := b.Create(beer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error in create beer"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}
