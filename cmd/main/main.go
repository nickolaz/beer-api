package main

import (
	"beer-api/beers/web"
	"beer-api/internal/database"
	"beer-api/internal/logs"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"path/filepath"
)

func main() {
	_ = logs.InitLogger()
	const userDB = "beer"
	const passDB = "beerApi"
	const hostDB = "localhost"
	const portDB = "5432"
	const nameDB = "db_beer"
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", userDB, passDB, hostDB, portDB, nameDB)
	client := database.NewSqlClient(connectionString)
	doMigrate(client, nameDB)
	createBeerHandler := web.NewCreateBeerHandler(client)
	mux := Routes(createBeerHandler)
	server := NewServer(mux)
	server.Run()
}

func doMigrate(client *database.PgClient, dbName string) {
	migrationsRootFolder, _ := filepath.Abs("../../migrations")
	driver, _ := postgres.WithInstance(client.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///"+migrationsRootFolder,
		dbName,
		driver,
	)
	if err != nil {
		logs.Log().Error(err.Error())
		return
	}
	err = m.Up()
	if err != nil && err.Error() != "no change" {
		logs.Log().Error(err.Error())
		return
	}
	if err.Error() == "no change" {
		logs.Log().Info("no migration needed")
	} else {
		current, _, _ := m.Version()
		logs.Log().Infof("current migration version %d ", current)
	}
}
