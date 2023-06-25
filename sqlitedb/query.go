package sqlitedb

import (
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	log "github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v4"
	"strings"
)

type Record struct {
	Id    string
	Hash  string
	Title string
	Dt    string
	Cat   string
	Size  uint64
	ExtId null.String
	Imdb  null.String
}

type SearchQuery struct {
	Query      string
	Categories []string
	ImdbId     null.String
	Season     null.Int
	Episode    null.Int
	Limit      uint
	Offset     uint
}

func Query(db *sql.DB, query *SearchQuery) ([]Record, error) {
	sqlQuery, err := createQuery(query)
	log.WithField("query", sqlQuery).Debug("SQL query")
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	var records []Record
	for rows.Next() {
		record := Record{}
		err = rows.Scan(&record.Id, &record.Hash, &record.Title, &record.Dt, &record.Cat, &record.Size, &record.ExtId, &record.Imdb)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func createQuery(query *SearchQuery) (string, error) {
	dialect := goqu.Dialect("sqlite3")
	var expressions []goqu.Expression
	if len(query.Query) > 0 {
		expressions = append(expressions, goqu.C("title").Like("%"+strings.ReplaceAll(strings.ToLower(query.Query), " ", ".")+".%"))
	}
	if len(query.Categories) > 0 {
		expressions = append(expressions, goqu.C("cat").In(query.Categories))
	}
	if query.ImdbId.Valid {
		imdbId := query.ImdbId.String
		if !strings.HasPrefix(imdbId, "tt") {
			imdbId = "tt" + imdbId
		}
		expressions = append(expressions, goqu.C("imdb").Eq(imdbId))
	}
	if query.Season.Valid && query.Episode.Valid {
		expressions = append(expressions, goqu.C("title").Like("%.S"+fmt.Sprintf("%02d", query.Season.Int64)+"E"+fmt.Sprintf("%02d", query.Episode.Int64)+".%"))
	} else if query.Season.Valid {
		expressions = append(expressions, goqu.C("title").Like("%.S"+fmt.Sprintf("%02d", query.Season.Int64)+"%"))
	} else if query.Episode.Valid {
		expressions = append(expressions, goqu.C("title").Like("%E"+fmt.Sprintf("%02d", query.Episode.Int64)+".%"))
	}
	ds := dialect.From("items").Where(expressions...).Order(goqu.C("dt").Desc()).Limit(query.Limit).Offset(query.Offset)
	sqlQuery, _, err := ds.ToSQL()
	return sqlQuery, err
}
