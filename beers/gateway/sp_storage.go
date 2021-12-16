package gateway

import (
	"beer-api/beers/models"
	"beer-api/internal/database"
	"beer-api/internal/logs"
	"database/sql"
	"fmt"
)

type BeerStorageGateway interface {
	Create(cmd *models.CreateBeerCMD) (*models.Beer, error)
	Get() (*[]models.Beer, error)
	BeerOfTheDay() (*models.Beer, error)
	GetDetails(beerId uint64) (*models.Beer, error)
}

type BeerStorage struct {
	*database.PgClient
}

func (b *BeerStorage) Create(cmd *models.CreateBeerCMD) (*models.Beer, error) {
	tx, err := b.PgClient.Begin()
	if err != nil {
		logs.Log().Error("Cannot create transaction")
		return nil, err
	}
	var id int64
	err = tx.QueryRow(
		`INSERT INTO beers (name, description , currency , brewery , country , price) VALUES ($1, $2, $3 , $4 , $5 , $6) RETURNING id`,
		cmd.Name, cmd.Description, cmd.Currency, cmd.Brewery, cmd.Country, cmd.Price).Scan(&id)
	if err != nil {
		logs.Log().Error("Cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}
	_ = tx.Commit()
	return &models.Beer{
		Id:          uint64(id),
		Name:        cmd.Name,
		Description: cmd.Description,
		Price:       cmd.Price,
		Country:     cmd.Country,
		Currency:    cmd.Currency,
		Brewery:     cmd.Brewery,
	}, nil
}

func (b *BeerStorage) Get() (*[]models.Beer, error) {
	tx, err := b.PgClient.Begin()
	if err != nil {
		logs.Log().Error("Cannot create transaction")
		return nil, err
	}
	res, err := tx.Query(`SELECT * from beers`)
	if err != nil {
		logs.Log().Error("Cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}
	var Beers []models.Beer
	for res.Next() {
		var beer models.Beer
		err = res.Scan(&beer.Id, &beer.Name, &beer.Description, &beer.Currency, &beer.Brewery, &beer.Country, &beer.Price)
		if err != nil {
			logs.Log().Error("Cannot get the beers ", err)
			_ = tx.Rollback()
			return nil, err
		}
		Beers = append(Beers, beer)
	}
	return &Beers, nil
}

func (b *BeerStorage) BeerOfTheDay() (*models.Beer, error) {
	tx, err := b.PgClient.Begin()
	if err != nil {
		logs.Log().Error("Cannot create transaction")
		return nil, err
	}
	res, err := tx.Query(`SELECT * FROM beers ORDER BY RANDOM() LIMIT 1`)
	if err != nil {
		logs.Log().Error("Cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}
	var Beer models.Beer
	for res.Next() {
		err = res.Scan(&Beer.Id, &Beer.Name, &Beer.Description, &Beer.Currency, &Beer.Brewery, &Beer.Country, &Beer.Price)
		if err != nil {
			logs.Log().Error("Cannot get the beers")
			_ = tx.Rollback()
			return nil, err
		}
	}
	return &Beer, nil
}

func (b *BeerStorage) GetDetails(beerId uint64) (*models.Beer, error) {
	tx, err := b.PgClient.Begin()
	if err != nil {
		logs.Log().Error("Cannot create transaction")
		return nil, err
	}
	queryString := fmt.Sprintf(`SELECT * FROM beers WHERE id = %d`, beerId)
	var Beer models.Beer
	err = tx.QueryRow(queryString).Scan(&Beer.Id, &Beer.Name, &Beer.Description, &Beer.Currency, &Beer.Brewery, &Beer.Country, &Beer.Price)
	if err == sql.ErrNoRows {
		logs.Log().Info("Cannot execute statement , id not found")
		_ = tx.Rollback()
		return nil, err
	}
	return &Beer, nil
}
