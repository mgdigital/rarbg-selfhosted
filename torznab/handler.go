package torznab

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v4"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func CreateHandler(db *sql.DB, trackers []string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = io.WriteString(w, "unable to parse form data")
			return
		}
		log.WithField("url", r.URL.String()).Debug("Received request")
		reqType := r.FormValue("t")
		switch reqType {
		case "caps":
			caps, err := Caps()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, err = io.WriteString(w, "internal server error")
				return
			}
			w.Header().Set("Content-Type", "application/xml")
			_, err = io.WriteString(w, caps)
			log.WithField("caps", caps).Debug("Handled caps request")
		case "search", "movie":
			q := r.FormValue("q")
			cat := r.FormValue("cat")
			var cats []int
			if cat != "" {
				catsStr := strings.Split(cat, ",")
				for i := range catsStr {
					id, err := strconv.Atoi(catsStr[i])
					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
						_, err = io.WriteString(w, "invalid category ID specified")
						return
					}
					cats = append(cats, id)
				}
			}
			queryImdbId := r.FormValue("imdbid")
			var imdbId null.String
			if queryImdbId != "" {
				imdbId.String = queryImdbId
				imdbId.Valid = true
			}
			result, err := Search(db, trackers, &SearchQuery{
				Query:    q,
				Cats:     cats,
				ImdbId:   imdbId,
				Attrs:    nil,
				Extended: false,
				Limit:    50,
				Offset:   0,
			})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, err = io.WriteString(w, "internal server error")
				return
			}
			_, err = io.WriteString(w, result)
			log.WithField("result", result).Debug("Handled search request")
		default:
			w.WriteHeader(http.StatusBadRequest)
			_, err = io.WriteString(w, "invalid request type")
			return
		}
		if err != nil {
			panic(err)
		}
	}
}
