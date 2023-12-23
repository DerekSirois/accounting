package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

func InitDb() error {
	DSN := os.Getenv("DSN")
	d, err := sqlx.Connect("postgres", DSN)
	if err != nil {
		return err
	}
	Db = d
	Db.MustExec(schema)

	return nil
}
