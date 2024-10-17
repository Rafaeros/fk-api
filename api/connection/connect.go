package connection

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func OpenConnection() (db *sql.DB, err error) {
	return sql.Open("sqlite3", "./fk.db")
}