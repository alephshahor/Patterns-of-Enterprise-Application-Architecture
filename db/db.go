package main

import (
	"github.com/go-pg/pg/v10"
	"os"
	"sync"
)

var db *pg.DB

func DB() *pg.DB {
	var lock = &sync.Mutex{}

	if db == nil {
		lock.Lock()
		defer lock.Unlock()

		db = pg.Connect(&pg.Options{
			Addr:     os.Getenv("PG_ADDR"),
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWD"),
			Database: os.Getenv("PG_DB"),
		})
	}

	return db
}
