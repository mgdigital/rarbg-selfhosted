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
			writeBadRequest(w, "unable to parse form data")
			return
		}
		reqType := r.FormValue("t")
		switch reqType {
		case "caps":
			caps, err := Caps()
			if err != nil {
				writeInternalError(w, err)
				return
			}
			w.Header().Set("Content-Type", "application/xml")
			writeResponseString(w, caps)
			log.WithField("caps", caps).Debug("handled caps request")
		case "search", "movie", "tv", "tvsearch":
			q := r.FormValue("q")
			cat := r.FormValue("cat")
			var cats []int
			if cat != "" {
				catsStr := strings.Split(cat, ",")
				for i := range catsStr {
					id, err := strconv.Atoi(catsStr[i])
					if err != nil {
						writeBadRequest(w, "invalid category ID specified")
						return
					}
					cats = append(cats, id)
				}
			}
			var limit = uint(50)
			var strQueryLimit = r.FormValue("limit")
			if strQueryLimit != "" {
				queryLimit, err := strconv.Atoi(strQueryLimit)
				if err != nil {
					writeBadRequest(w, "invalid limit specified")
					return
				}
				limit = uint(queryLimit)
			}
			var offset = uint(0)
			var strQueryOffset = r.FormValue("offset")
			if strQueryOffset != "" {
				queryOffset, err := strconv.Atoi(strQueryOffset)
				if err != nil {
					writeBadRequest(w, "invalid offset specified")
					return
				}
				offset = uint(queryOffset)
			}
			var imdbId null.String
			queryImdbId := r.FormValue("imdbid")
			if queryImdbId != "" {
				imdbId.String = queryImdbId
				imdbId.Valid = true
			}
			var season null.Int
			strQuerySeason := r.FormValue("season")
			if strQuerySeason != "" {
				querySeason, err := strconv.ParseInt(strQuerySeason, 10, 64)
				if err != nil {
					writeBadRequest(w, "invalid season specified")
					return
				}
				season.Int64 = querySeason
				season.Valid = true
			}
			var episode null.Int
			strQueryEpisode := r.FormValue("ep")
			if strQueryEpisode != "" {
				queryEpisode, err := strconv.ParseInt(strQueryEpisode, 10, 64)
				if err != nil {
					writeBadRequest(w, "invalid episode specified")
					return
				}
				episode.Int64 = queryEpisode
				episode.Valid = true
			}
			searchQuery := &SearchQuery{
				Query:    q,
				Cats:     cats,
				ImdbId:   imdbId,
				Season:   season,
				Episode:  episode,
				Attrs:    nil,
				Extended: false,
				Limit:    limit,
				Offset:   offset,
			}
			result, err := Search(db, trackers, searchQuery)
			if err != nil {
				writeInternalError(w, err)
				return
			}
			writeResponseString(w, result)
		default:
			writeBadRequest(w, "invalid request type")
			return
		}
	}
}

func writeBadRequest(w http.ResponseWriter, message string) {
	log.WithField("message", message).Warn("bad request")
	writeErrorResponse(w, http.StatusBadRequest, message)
}

func writeInternalError(w http.ResponseWriter, err error) {
	log.WithField("error", err).Error("internal server error")
	writeErrorResponse(w, http.StatusInternalServerError, "internal server error")
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	writeResponseString(w, message)
}

func writeResponseString(w http.ResponseWriter, str string) {
	_, err := io.WriteString(w, str)
	if err != nil {
		panic(err)
	}
}
