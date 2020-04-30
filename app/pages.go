package app

import (
	"encoding/json"

	"github.com/dgraph-io/badger"
	"github.com/gomarkdown/markdown"
)

type Page struct {
	Title string
	Body  string
}

func (p *Page) Render() string {
	md := []byte(p.Body)
	output := markdown.ToHTML(md, nil, nil)
	return string(output)
}

func (p *Page) save() error {
	return db.Update(func(txn *badger.Txn) error {
		serialized, err := json.Marshal(p)
		if err != nil {
			println(err.Error())
			return err
		}
		return txn.Set([]byte(p.Title), serialized)
	})
}

func loadPage(title string) (*Page, error) {
	var page *Page
	err := db.View(func(txn *badger.Txn) error {
		val, err := txn.Get([]byte(title))
		if err != nil {
			return err
		}
		err = val.Value(func(val []byte) error {
			return json.Unmarshal(val, &page)
		})
		return err
	})
	return page, err
}
