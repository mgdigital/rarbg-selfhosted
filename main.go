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
	if os.Getenv("DEBUG") == "1" {
		log.SetLevel(log.DebugLevel)
	}
	dbPath := os.Getenv("PATH_SQLITE_DB")
	if dbPath == "" {
		dbPath = "./rarbg_db.sqlite"
	}
	trackersPath := os.Getenv("PATH_TRACKERS")
	if trackersPath == "" {
		trackersPath = "./trackers.txt"
	}
	db, err := sqlitedb.Open(dbPath)
	if err != nil {
		panic(err)
	}
	trackers, err := magnet.GetTrackers(trackersPath)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/torznab/", torznab.CreateHandler(db, trackers))
	log.Info("starting server on port 3333")
	err = http.ListenAndServe(":3333", logRequest(http.DefaultServeMux))
	if err != nil {
		panic(err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
		}).Debug("http request")
		handler.ServeHTTP(w, r)
	})
}
