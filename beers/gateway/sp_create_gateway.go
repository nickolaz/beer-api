package gateway

import (
	"beer-api/beers/models"
	"beer-api/internal/database"
)

type BeerCreateGateway interface {
	Create(cmd *models.CreateBeerCMD) (*models.Beer, error)
}

type BeerCreateGtw struct {
	BeerStorageGateway
}

func NewBeerCreateGateway(client *database.PgClient) BeerCreateGateway {
	return &BeerCreateGtw{&BeerStorage{client}}
}
