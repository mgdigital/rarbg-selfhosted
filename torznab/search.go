package torznab

import (
	"bytes"
	"database/sql"
	xw "github.com/shabbyrobe/xmlwriter"
	"gopkg.in/guregu/null.v4"
	"mgdigital/rarbg-selfhosted/magnet"
	"mgdigital/rarbg-selfhosted/sqlitedb"
	"strconv"
	"strings"
	"time"
)

func Search(db *sql.DB, trackers []string, query *SearchQuery) (string, error) {
	result, err := doSearch(db, trackers, query)
	if err != nil {
		return "", err
	}
	var contentNode []xw.Writable
	contentNode = append(contentNode, xw.Elem{
		Name: "title",
		Content: []xw.Writable{
			xw.Text("RARBG SQLite"),
		},
	})
	for i := range result {
		contentNode = append(contentNode, xw.Elem{
			Name: "item",
			Content: []xw.Writable{
				xw.Elem{
					Name: "title",
					Content: []xw.Writable{
						xw.Text(result[i].Title),
					},
				},
				xw.Elem{
					Name: "guid",
					Content: []xw.Writable{
						xw.Text(result[i].Guid),
					},
				},
				xw.Elem{
					Name: "type",
					Content: []xw.Writable{
						xw.Text("public"),
					},
				},
				xw.Elem{
					Name: "pubDate",
					Content: []xw.Writable{
						xw.Text(result[i].PubDate.String()),
					},
				},
				xw.Elem{
					Name: "size",
					Content: []xw.Writable{
						xw.Text(strconv.FormatUint(result[i].Size, 10)),
					},
				},
				xw.Elem{
					Name: "category",
					Content: []xw.Writable{
						xw.Text(result[i].Category),
					},
				},
				xw.Elem{
					Name: "enclosure",
					Attrs: []xw.Attr{
						{Name: "url", Value: result[i].MagnetLink},
						{Name: "length", Value: strconv.FormatUint(result[i].Size, 10)},
						{Name: "type", Value: "application/x-bittorrent;x-scheme-handler/magnet"},
					},
				},
				xw.Elem{
					Name: "torznab:attr",
					Attrs: []xw.Attr{
						{Name: "name", Value: "category"},
						{Name: "value", Value: strconv.Itoa(result[i].CategoryId)},
					},
				},
			},
		})
	}
	b := &bytes.Buffer{}
	w := xw.Open(b)
	ec := &xw.ErrCollector{}
	defer ec.Panic()
	ec.Do(
		w.Start(xw.Doc{}),
		w.Start(xw.Elem{
			Name: "rss",
			Attrs: []xw.Attr{
				{Name: "version", Value: "2.0"},
				{Name: "xmlns:atom", Value: "http://www.w3.org/2005/Atom"},
				{Name: "xmlns:torznab", Value: "http://torznab.com/schemas/2015/feed"},
			},
			Content: []xw.Writable{
				xw.Elem{
					Name:    "channel",
					Content: contentNode,
				},
			},
		}),
		w.EndAllFlush(),
	)
	if ec.Err != nil {
		return "", ec.Err
	}
	return b.String(), nil

}

type SearchQuery struct {
	Query    string
	Cats     []int
	ImdbId   null.String
	Attrs    []string
	Extended bool
	Limit    uint
	Offset   uint
}

type SearchResultItem struct {
	Title       string
	Guid        string
	PubDate     time.Time
	Category    string
	CategoryId  int
	Size        uint64
	Description string
	MagnetLink  string
}

func doSearch(db *sql.DB, trackers []string, query *SearchQuery) ([]SearchResultItem, error) {
	var queryCategoryNames []string
	for catId := range query.Cats {
		cats := IdToCategories(catId)
		for i := range cats {
			queryCategoryNames = append(queryCategoryNames, cats[i])
		}
	}
	dbResult, err := sqlitedb.Query(db, &sqlitedb.SearchQuery{
		Query:      query.Query,
		Categories: queryCategoryNames,
		ImdbId:     query.ImdbId,
		Limit:      query.Limit,
		Offset:     query.Offset,
	})
	if err != nil {
		return nil, err
	}
	var resultItems []SearchResultItem
	for i := range dbResult {
		item, err := transformResultItem(dbResult[i], trackers)
		if err != nil {
			return nil, err
		}
		resultItems = append(resultItems, *item)
	}
	return resultItems, nil
}

func transformResultItem(record sqlitedb.Record, trackers []string) (*SearchResultItem, error) {
	pubDate, err := time.Parse("2006-01-02 15:04:05", record.Dt)
	if err != nil {
		return nil, err
	}
	return &SearchResultItem{
		Title:       strings.ReplaceAll(record.Title, ".", " "),
		Guid:        record.Id,
		PubDate:     pubDate,
		Category:    record.Cat,
		CategoryId:  CategoryToId(record.Cat),
		Size:        record.Size,
		Description: record.Title,
		MagnetLink:  magnet.CreateMagnetLink(record.Hash, record.Title, trackers),
	}, nil
}
