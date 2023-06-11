package main

import (
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"mgdigital/rarbg-selfhosted/magnet"
	"mgdigital/rarbg-selfhosted/sqlitedb"
	"mgdigital/rarbg-selfhosted/torznab"
	"net/http"
	"os"
)

func main() {
	dbPath := os.Getenv("PATH_SQLITE_DB")
	trackersPath := os.Getenv("PATH_TRACKERS")
	log.Info("Loading SQLite DB...")
	db, err := sqlitedb.Open(dbPath)
	if err != nil {
		panic(err)
	}
	log.Info("Loading trackers file...")
	trackers, err := magnet.GetTrackers(trackersPath)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/torznab/", torznab.CreateHandler(db, trackers))
	log.Info("Starting server...")
	err = http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}
