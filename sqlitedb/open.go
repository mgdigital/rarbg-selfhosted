package sqlitedb

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"os"
)

func Open(dbPath string) (*sql.DB, error) {
	log.Info("opening database file")
	info, err := os.Stat(dbPath)
	if err != nil {
		log.WithField("error", err).Error("failed to stat database file")
		return nil, err
	}
	if info.IsDir() {
		err = errors.New("failed to open database: is a directory")
		log.WithField("error", err).Error(err.Error())
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.WithField("error", err).Error("failed to open database")
		return nil, err
	}
	log.Debug("performing database sanity check")
	_, err = db.Query("select id, hash, title, dt, cat, size, ext_id, imdb from items limit 1")
	if err != nil {
		log.WithField("error", err).Error("database sanity check failed")
		return nil, err
	}
	return db, nil
}
