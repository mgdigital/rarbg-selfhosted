package torznab

import (
	"bytes"
	xw "github.com/shabbyrobe/xmlwriter"
	"strconv"
)

func Caps() (string, error) {
	var categoryNodes []xw.Writable
	for name, id := range categoryMap {
		categoryNodes = append(categoryNodes, xw.Elem{
			Name: "category",
			Attrs: []xw.Attr{
				{Name: "id", Value: strconv.Itoa(id)},
				{Name: "name", Value: name},
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
			Name: "caps",
			Content: []xw.Writable{
				xw.Elem{
					Name: "server", Attrs: []xw.Attr{
						{Name: "version", Value: "1.1"},
						{Name: "title", Value: "RARBG SQLite DB"},
					},
				},
				xw.Elem{
					Name: "limits", Attrs: []xw.Attr{
						{Name: "max", Value: "100"},
						{Name: "default", Value: "50"},
					},
				},
				xw.Elem{
					Name: "searching",
					Content: []xw.Writable{
						xw.Elem{
							Name: "search", Attrs: []xw.Attr{
								{Name: "available", Value: "yes"},
								{Name: "supportedParams", Value: "q"},
							},
						},
						xw.Elem{
							Name: "tv-search", Attrs: []xw.Attr{
								{Name: "available", Value: "yes"},
								{Name: "supportedParams", Value: "q,imdbid,season,ep"},
							},
						},
						xw.Elem{
							Name: "movie-search", Attrs: []xw.Attr{
								{Name: "available", Value: "yes"},
								{Name: "supportedParams", Value: "q,imdbid"},
							},
						},
						xw.Elem{
							Name: "audio-search", Attrs: []xw.Attr{
								{Name: "available", Value: "no"},
								{Name: "supportedParams", Value: "q"},
							},
						},
						xw.Elem{
							Name: "book-search", Attrs: []xw.Attr{
								{Name: "available", Value: "no"},
								{Name: "supportedParams", Value: "q"},
							},
						},
					},
				},
				xw.Elem{
					Name: "categories", Content: categoryNodes,
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
