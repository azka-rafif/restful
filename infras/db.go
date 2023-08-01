package infras

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func GetDB() (*sqlx.DB, error) {

	db, err := sqlx.Open("sqlite3", "infras/new.db")
	return db, err
}
