package database

import (
	"beer-api/internal/logs"
	"database/sql"
	_ "github.com/lib/pq"
)

type PgClient struct {
	*sql.DB
}

func NewSqlClient(source string) *PgClient {
	db, err := sql.Open("postgres", source)
	if err != nil {
		logs.Log().Errorf("cannot create db connection: %s", err.Error())
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		logs.Log().Errorf("cannot connect to db: %s", err.Error())
	}
	return &PgClient{db}
}
