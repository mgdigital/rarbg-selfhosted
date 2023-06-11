package sqlitedb

import (
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
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
	Limit      uint
	Offset     uint
}

func Query(db *sql.DB, query *SearchQuery) ([]Record, error) {
	dialect := goqu.Dialect("sqlite3")
	var expressions []goqu.Expression
	if len(query.Query) > 0 {
		expressions = append(expressions, goqu.C("title").Like("%"+strings.ReplaceAll(strings.ToLower(query.Query), " ", ".")+".%"))
	}
	if len(query.Categories) > 0 {
		expressions = append(expressions, goqu.C("cat").In(query.Categories))
	}
	if query.ImdbId.Valid {
		expressions = append(expressions, goqu.C("imdb").Eq(query.ImdbId.String))
	}
	ds := dialect.From("items").Where(expressions...).Order(goqu.C("dt").Desc()).Limit(query.Limit).Offset(query.Offset)
	sqlQuery, params, err := ds.ToSQL()
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v", params)
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
