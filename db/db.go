package db

import (
	"os"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB

func DB() *pg.DB {
	if db == nil {
		db = pg.Connect(&pg.Options{
			Addr:     os.Getenv("PG_ADDR"),
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWD"),
			Database: os.Getenv("PG_DB"),
		})
	}
	return db
}
