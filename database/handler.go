package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS user(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(24) NOT NULL,
    password VARCHAR(65) NOT NULL
);
`

func NewConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "db.sql")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
