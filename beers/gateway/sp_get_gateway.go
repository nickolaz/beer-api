package gateway

import (
	"beer-api/beers/models"
	"beer-api/internal/database"
)

type BeerGetGateway interface {
	Get() (*[]models.Beer, error)
	BeerOfTheDay() (*models.Beer, error)
	GetDetails(beerId uint64) (*models.Beer, error)
}

type BeerGetGtw struct {
	BeerStorageGateway
}

func GetBeerGateway(client *database.PgClient) BeerGetGateway {
	return &BeerGetGtw{&BeerStorage{client}}
}
