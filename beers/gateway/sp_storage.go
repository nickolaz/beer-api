package gateway

import (
	"beer-api/beers/models"
	"beer-api/internal/database"
	"beer-api/internal/logs"
)

type BeerStorageGateway interface {
	Create(cmd *models.CreateBeerCMD) (*models.Beer, error)
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
	res, err := tx.Exec(`INSERT INTO beers (name, description, price ) VALUES (?, ?, ? )`, cmd.Name, cmd.Description, cmd.Price)
	if err != nil {
		logs.Log().Error("Cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		logs.Log().Error("Cannot get last insert id")
		_ = tx.Rollback()
		return nil, err
	}
	_ = tx.Commit()
	return &models.Beer{
		Id:          uint64(int(id)),
		Name:        cmd.Name,
		Description: cmd.Description,
		Price:       cmd.Price,
	}, nil
}
