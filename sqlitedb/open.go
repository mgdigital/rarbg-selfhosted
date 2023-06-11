package sqlitedb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func Open(dbPath string) (*sql.DB, error) {
	_, err := os.Stat(dbPath)
	if err != nil {
		return nil, err
	}
	return sql.Open("sqlite3", dbPath)
}
