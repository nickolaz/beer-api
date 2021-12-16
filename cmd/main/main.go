package main

import (
	"beer-api/beers/web"
	_ "beer-api/docs"
	"beer-api/internal/database"
	"beer-api/internal/logs"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"os"
	"path/filepath"
)

// @title           Falabella Challenge Api
// @version         1.0
// @description     Api for a Falabella Challenge.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Nicolas Ayala
// @contact.url    https://py.linkedin.com/in/nicolas-ayala-koy
// @contact.email  nicoayalakoy@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	_ = logs.InitLogger()
	userDB, passDB, nameDB, portDB, hostDB :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_HOST")
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", userDB, passDB, hostDB, portDB, nameDB)
	client := database.NewSqlClient(connectionString)
	doMigrate(client, nameDB)
	createBeerHandler := web.NewCreateBeerHandler(client)
	getBeerHandler := web.NewGetBeerHandler(client)
	mux := Routes(createBeerHandler, getBeerHandler)
	server := NewServer(mux)
	server.Run()
}

func doMigrate(client *database.PgClient, dbName string) {
	migrationsRootFolder, _ := filepath.Abs("migrations")
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
	if err != nil && err.Error() == "no change" {
		logs.Log().Info("no migration needed")
	} else {
		current, _, _ := m.Version()
		logs.Log().Infof("current migration version %d ", current)
	}
}
